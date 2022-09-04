package pingengine

import (
	"networkmonitor/net/pinger"
	"networkmonitor/timer"
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
