package pingengine

import (
	"networkmonitor/core/net/pinger"
	"networkmonitor/core/timer"
)

type PingResultHandler interface {
	OnPingResultHandle(pinger.Stats)
}

type PingTimerHandlerFactory interface {
	CreatePingTimerHandler(pingResultHandler PingResultHandler, ipaddress string, pingCount int) timer.TimerHandler
}

type Engine interface {
	RegisterIpAddress(ipAddress string) bool
	Run() error
}
