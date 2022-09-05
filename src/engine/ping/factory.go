package pingengine

import (
	db "networkmonitor/db/kv"
	"networkmonitor/net/pinger"
	"networkmonitor/timer"
)

// interface compliance
var _ PingTimerHandlerFactory = &pingTimerHandlerFactory{}

func MakePingTimerHandlerFactory() PingTimerHandlerFactory {
	return &pingTimerHandlerFactory{}
}

type pingTimerHandlerFactory struct{}

func (p *pingTimerHandlerFactory) CreatePingTimerHandler(pingResultHandler PingResultHandler, db db.KvDb, pingCount int) timer.TimerHandler {
	return MakePingTimerHandler(pingResultHandler, pinger.MakePinger(), db, pingCount)
}
