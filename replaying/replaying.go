package replaying

import (
	"context"
	"encoding/json"
	"github.com/v2pro/koala/recording"
	"github.com/v2pro/plz/countlog"
	"net"
	"strconv"
	"strings"
)

type ReplayingSession struct {
	SessionId         string
	CallFromInbound   *recording.CallFromInbound
	ReturnInbound     *recording.ReturnInbound
	CallOutbounds     []*recording.CallOutbound
	RedirectDirs      map[string]string
	MockFiles         map[string][]byte
	TracePaths        []string
	actionCollector   chan ReplayedAction
	lastMaxScoreIndex int
}

func NewReplayingSession() *ReplayingSession {
	return &ReplayingSession{
		actionCollector:   make(chan ReplayedAction, 40960),
		lastMaxScoreIndex: -1,
	}
}

func (replayingSession *ReplayingSession) collectAction(ctx context.Context, action ReplayedAction) {
	select {
	case replayingSession.actionCollector <- action:
	default:
		countlog.Error("event!replaying.ActionCollector is full", "ctx", ctx)
	}
}

func (replayingSession *ReplayingSession) CallOutbound(ctx context.Context, callOutbound *CallOutbound) {
	replayingSession.collectAction(ctx, callOutbound)
}

func (replayingSession *ReplayingSession) CallFunction(ctx context.Context, content []byte) {
	callFunction := &CallFunction{}
	err := json.Unmarshal(content, callFunction)
	if err != nil {
		countlog.Error("event!replaying.unmarshal CallFunction failed", "err", err, "content", content)
		return
	}
	callFunction.ActionType = "CallFunction"
	callFunction.OccurredAt, _ = strconv.ParseInt(callFunction.ActionId, 10, 64)
	replayingSession.collectAction(ctx, callFunction)
}

func (replayingSession *ReplayingSession) ReturnFunction(ctx context.Context, content []byte) {
	returnFunction := &ReturnFunction{}
	err := json.Unmarshal(content, returnFunction)
	if err != nil {
		countlog.Error("event!replaying.unmarshal ReturnFunction failed", "err", err, "content", content)
		return
	}
	returnFunction.replayedAction = newReplayedAction("ReturnFunction")
	replayingSession.collectAction(ctx, returnFunction)
}

func (replayingSession *ReplayingSession) AppendFile(ctx context.Context, content []byte, fileName string) {
	if replayingSession == nil {
		return
	}
	appendFile := &AppendFile{
		replayedAction: newReplayedAction("AppendFile"),
		FileName:       fileName,
		Content:        append([]byte(nil), content...),
	}
	replayingSession.collectAction(ctx, appendFile)
}

func (replayingSession *ReplayingSession) SendUDP(ctx context.Context, content []byte, peer net.UDPAddr) {
	if replayingSession == nil {
		return
	}
	sendUdp := &SendUDP{
		replayedAction: newReplayedAction("SendUDP"),
		Peer:           peer,
		Content:        append([]byte(nil), content...),
	}
	replayingSession.collectAction(ctx, sendUdp)
}

func (replayingSession *ReplayingSession) Finish(response []byte) *ReplayedSession {
	replayedSession := &ReplayedSession{
		SessionId: replayingSession.SessionId,
		CallFromInbound: &CallFromInbound{
			replayedAction:      newReplayedAction("CallFromInbound"),
			OriginalRequestTime: replayingSession.CallFromInbound.OccurredAt,
			OriginalRequest:     replayingSession.CallFromInbound.Request,
		},
	}
	replayedSession.ReturnInbound = &ReturnInbound{
		replayedAction:   newReplayedAction("ReturnInbound"),
		OriginalResponse: replayingSession.ReturnInbound.Response,
		Response:         response,
	}
	done := false
	for !done {
		select {
		case action := <-replayingSession.actionCollector:
			replayedSession.Actions = append(replayedSession.Actions, action)
		default:
			done = true
		}
	}
	replayedSession.Actions = append(replayedSession.Actions, replayedSession.ReturnInbound)
	return replayedSession
}

func (replayingSession *ReplayingSession) ShouldTraceFile(fileName string) bool {
	for _, tracePath := range replayingSession.TracePaths {
		if strings.HasPrefix(fileName, tracePath) {
			return true
		}
	}
	return false
}
