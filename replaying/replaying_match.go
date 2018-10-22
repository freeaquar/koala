package replaying

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"sync"

	"github.com/v2pro/koala/recording"
	"github.com/v2pro/plz/countlog"
)

var expect100 = []byte("Expect: 100-continue")

// outbound level's last matched index
// sessionId map outboundsLastMatchedIndex
var globalLastMatchedIndex = map[string]int{}
var globalLastMatchedIndexMutex = &sync.Mutex{}

func (replayingSession *ReplayingSession) MatchOutboundTalk(
	ctx context.Context, connLastMatchedIndex int, request []byte) (int, float64, *recording.CallOutbound) {
	//return replayingSession.MinHashMatchOutboundTalk(ctx, connLastMatchedIndex, request)

	unit := 16
	chunks := cutToChunks(request, unit)

	outboundsRevealData := replayingSession.revealSessions()
	requestRevealData := replayingSession.revealOneSession(request)
	reqCandidates := replayingSession.loadKeys()
	outboundsLen := len(replayingSession.CallOutbounds)

	scores := make([]int, outboundsLen)
	reqExpect100 := bytes.Contains(request, expect100)

	for i, callOutbound := range replayingSession.CallOutbounds {
		if reqExpect100 != bytes.Contains(callOutbound.Request, expect100) {
			scores[i] = math.MinInt64
		}
		if !outboundsRevealData[i].Handler.PreMatch(requestRevealData, outboundsRevealData[i]) {
			scores[i] = math.MinInt64
		}
	}

	beginMatchIndex := getBeginMatchIndex(replayingSession.SessionId, connLastMatchedIndex)
	endMatchIndex := outboundsLen - 1

	maxScoreIndex, mark := replayingSession.matchTalk(ctx, chunks, reqCandidates, beginMatchIndex, endMatchIndex, connLastMatchedIndex, scores)
	if maxScoreIndex == -1 && beginMatchIndex > 0 {
		// try from 0 to beginIndex - 1
		maxScoreIndex, mark = replayingSession.matchTalk(ctx, chunks, reqCandidates, 0, beginMatchIndex-1, connLastMatchedIndex, scores)
	}
	if maxScoreIndex == -1 {
		return -1, 0, nil
	}

	if beginMatchIndex < maxScoreIndex {
		setGlobalLastMatchedIndex(replayingSession.SessionId, maxScoreIndex)
	}
	return maxScoreIndex, mark, replayingSession.CallOutbounds[maxScoreIndex]
}

func (replayingSession *ReplayingSession) loadKeys() [][]byte {
	keys := make([][]byte, len(replayingSession.CallOutbounds))
	for i, entry := range replayingSession.CallOutbounds {
		keys[i] = entry.Request
	}
	return keys
}

func cutToChunks(key []byte, unit int) [][]byte {
	chunks := [][]byte{}
	if len(key) > 256 {
		offset := 0
		for {
			strikeStart, strikeLen := findReadableChunk(key[offset:])
			if strikeStart == -1 {
				break
			}
			if strikeLen > 8 {
				firstChunkLen := strikeLen
				if firstChunkLen > 16 {
					firstChunkLen = 16
				}
				chunks = append(chunks, key[offset+strikeStart:offset+strikeStart+firstChunkLen])
				key = key[offset+strikeStart+firstChunkLen:]
				break
			}
			offset += strikeStart + strikeLen
		}
	}
	chunkCount := len(key) / unit
	for i := 0; i < chunkCount; i++ {
		chunks = append(chunks, key[i*unit:(i+1)*unit])
	}
	lastChunk := key[chunkCount*unit:]
	if len(lastChunk) > 0 {
		chunks = append(chunks, lastChunk)
	}
	return chunks
}

// findReadableChunk returns: the starting index of the trunk, length of the trunk
func findReadableChunk(key []byte) (int, int) {
	start := bytes.IndexFunc(key, func(r rune) bool {
		return r > 31 && r < 127
	})
	if start == -1 {
		return -1, -1
	}
	end := bytes.IndexFunc(key[start:], func(r rune) bool {
		return r <= 31 || r >= 127
	})
	if end == -1 {
		return start, len(key) - start
	}
	return start, end
}

func getGlobalLastMatchedIndex(sessionId string) int {
	globalLastMatchedIndexMutex.Lock()
	defer globalLastMatchedIndexMutex.Unlock()
	if index, ok := globalLastMatchedIndex[sessionId]; ok {
		return index
	}
	return -1
}

func setGlobalLastMatchedIndex(sessionId string, outboundsLastMatchIndex int) {
	globalLastMatchedIndexMutex.Lock()
	defer globalLastMatchedIndexMutex.Unlock()
	globalLastMatchedIndex[sessionId] = outboundsLastMatchIndex
}

func getBeginMatchIndex(sessionId string, connLastMatchedIndex int) int {
	beginIndex := connLastMatchedIndex
	if connLastMatchedIndex == -1 {
		beginIndex = getGlobalLastMatchedIndex(sessionId)
	}
	return beginIndex + 1
}

func (replayingSession *ReplayingSession) matchTalk(ctx context.Context, reqChunks [][]byte, reqCandidates [][]byte, beginIndex int, endIndex int,
	connLastMatchedIndex int, scores []int) (int, float64) {
	maxScore := 0
	maxScoreIndex := -1
	for chunkIndex, chunk := range reqChunks {
		for j, reqCandidate := range reqCandidates {
			if j < beginIndex || len(reqCandidate) < len(chunk) {
				continue
			}
			if j > endIndex {
				break
			}
			pos := bytes.Index(reqCandidate, chunk)
			if pos >= 0 {
				reqCandidates[j] = reqCandidate[pos:]
				if chunkIndex == 0 && connLastMatchedIndex == -1 {
					scores[j] += len(reqChunks) // first chunk has more weight
				} else {
					scores[j]++
				}
				if scores[j] > maxScore {
					maxScore = scores[j]
					maxScoreIndex = j
				}
			}
		}
	}

	// 多个 maxScore，优先从上一次成功匹配的 Index 之后开始，取第一个 maxScore
	for j, score := range scores {
		if score == maxScore && replayingSession.lastMaxScoreIndex < j {
			maxScoreIndex = j
			break
		}
	}

	countlog.Trace("event!replaying.talks_scored",
		"ctx", ctx,
		"connLastMatchedIndex", connLastMatchedIndex,
		//"lastMaxScoreIndex", lastMatchedIndex,
		"beginMatchIndex", beginIndex,
		"endMatchIndex", endIndex,
		"maxScoreIndex", maxScoreIndex,
		"maxScore", maxScore,
		"totalScore", len(reqChunks),
		"scores", func() string {
			return fmt.Sprintf("%v", scores)
		})

	if maxScore == 0 {
		return -1, 0
	}
	mark := float64(maxScore) / float64(len(reqChunks))
	if connLastMatchedIndex != -1 {
		// not starting from beginning, should have minimal score
		if mark < 0.85 {
			return -1, 0
		}
	} else {
		if mark < 0.1 {
			return -1, 0
		}
	}

	if maxScoreIndex > replayingSession.lastMaxScoreIndex {
		replayingSession.lastMaxScoreIndex = maxScoreIndex
	}

	return maxScoreIndex, mark
}
