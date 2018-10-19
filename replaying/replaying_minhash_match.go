package replaying

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dgryski/go-metro"
	"github.com/dgryski/go-minhash"
	"github.com/dgryski/go-spooky"
	"github.com/v2pro/koala/recording"
	"github.com/v2pro/plz/countlog"
)

var globalMinHash = map[string][]*minhash.MinWise{}
var globalMinHashMutex = &sync.Mutex{}

func (replayingSession *ReplayingSession) MinHashMatchOutboundTalk(
	ctx context.Context, connLastMatchedIndex int, request []byte) (int, float64, *recording.CallOutbound) {
	maxScoreIndex := -1
	maxScore := float64(0)
	scores := make([]float64, len(replayingSession.CallOutbounds))
	outboundsHash := getReplayingSessionMinHash(replayingSession)
	requestHash := NewMinHash(request)

	beginMatchIndex := getBeginMatchIndex(replayingSession.SessionId, connLastMatchedIndex)
	for i := range replayingSession.CallOutbounds {
		if i < beginMatchIndex {
			continue
		}
		scores[i] = outboundsHash[i].Similarity(requestHash)
		if scores[i] > maxScore {
			maxScore = scores[i]
			maxScoreIndex = i
		}
		if scores[i] >= 0.96 {
			break
		}
	}
	for i, score := range scores {
		if score == maxScore && i >= beginMatchIndex {
			maxScoreIndex = i
			break
		}
	}

	if beginMatchIndex < maxScoreIndex {
		setGlobalLastMatchedIndex(replayingSession.SessionId, maxScoreIndex)
	}

	countlog.Trace("event!replaying.min_hash_talks_scored", "ctx", ctx, "connLastMatchedIndex", connLastMatchedIndex,
		"beginMatchIndex", beginMatchIndex, "maxScoreIndex", maxScoreIndex,
		"maxScore", maxScore,
		"scores", func() interface{} {
			return fmt.Sprintf("%v", scores)
		})
	if maxScore == 0 {
		return -1, 0, nil
	}
	return maxScoreIndex, scores[maxScoreIndex], replayingSession.CallOutbounds[maxScoreIndex]
}

//func hashMatchTalk(ctx context.Context, requestHash *minhash.MinWise, outboundsHash *minhash.MinWise,
//	beginIndex int, endIndex int, scores []float64) {
//	for i, _ := range replayingSession.CallOutbounds {
//		if i < beginIndex {
//			continue
//		}
//		if i > endIndex {
//			break
//		}
//		scores[i] = outboundsHash[i].Similarity(requestHash)
//		if scores[i] > maxScore {
//			maxScore = scores[i]
//			maxScoreIndex = i
//		}
//		if scores[i] >= 0.96 {
//			break
//		}
//	}
//}
func getReplayingSessionMinHash(replayingSession *ReplayingSession) []*minhash.MinWise {
	globalMinHashMutex.Lock()
	defer globalMinHashMutex.Unlock()
	outboundsHash := globalMinHash[replayingSession.SessionId]
	if outboundsHash == nil {
		begin := time.Now()
		outboundsHash = make([]*minhash.MinWise, len(replayingSession.CallOutbounds))
		for i, callOutbound := range replayingSession.CallOutbounds {
			outboundsHash[i] = NewMinHash(callOutbound.Request)
		}
		globalMinHash[replayingSession.SessionId] = outboundsHash
		elapsed := time.Since(begin)
		countlog.Trace("event!replaying.build_min_hash", "spendTime", elapsed)
	}
	return outboundsHash
}

func stopFieldsFunc(c rune) bool {
	return c == '\n' || c == '\r' || c == '&' || c == ',' || c == ':' || c == ' ' || c == '='
}

func mhash(b []byte) uint64 {
	return metro.Hash64(b, 0)
}

func NewMinHash(data []byte) *minhash.MinWise {
	hash := minhash.NewMinWise(spooky.Hash64, mhash, 64)
	chunks := cutToMinHashChunks(data)
	for _, chunk := range chunks {
		hash.Push(chunk)
	}
	return hash
}

func cutToMinHashChunks(key []byte) [][]byte {
	chunks := [][]byte{}
	offset := 0
	for {
		strikeStart, strikeLen := findReadableChunk(key[offset:])
		if strikeStart == -1 {
			break
		}
		for i := offset; i < offset+strikeStart; i += 1 {
			chunks = append(chunks, key[i:i+1]) // unreadable char
		}
		readableChunks := bytes.FieldsFunc(key[offset+strikeStart:offset+strikeStart+strikeLen], stopFieldsFunc)
		for _, readableChunk := range readableChunks {
			chunks = append(chunks, readableChunk) // readable chunk
		}
		offset += strikeStart + strikeLen
	}

	keyLen := len(key)
	for i := offset; i < keyLen; i += 1 {
		chunks = append(chunks, key[i:i+1]) // unreadable char
	}
	return chunks
}
