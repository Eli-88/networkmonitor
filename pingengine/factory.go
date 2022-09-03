package pingengine

import (
	"networkmonitor/core/net/pinger"
	"networkmonitor/core/timer"
)

// interface compliance
var _ PingTimerHandlerFactory = &pingTimerHandlerFactory{}

func MakePingTimerHandlerFactory() PingTimerHandlerFactory {
	return &pingTimerHandlerFactory{}
}

type pingTimerHandlerFactory struct{}

func (p *pingTimerHandlerFactory) CreatePingTimerHandler(pingResultHandler PingResultHandler, ipaddress string, pingCount int) timer.TimerHandler {
	return MakePingTimerHandler(pingResultHandler, pinger.MakePinger(), ipaddress, pingCount)
}
