package pingengine

import (
	db "networkmonitor/db/kv"
	"networkmonitor/net/pinger"
	"networkmonitor/timer"
)

type PingResultHandler interface {
	OnPingResultHandle(pinger.Stats)
}

type PingTimerHandlerFactory interface {
	CreatePingTimerHandler(pingResultHandler PingResultHandler, db db.KvDb, pingCount int) timer.TimerHandler
}

type Engine interface {
	RegisterIpAddress(ipAddress string)
	Run() error
}
