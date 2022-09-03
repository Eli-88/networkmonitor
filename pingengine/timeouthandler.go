package pingengine

import (
	"networkmonitor/core/logger"
	"networkmonitor/core/net/pinger"
	"networkmonitor/core/timer"
)

// interface compliance
var _ timer.TimerHandler = &pingHandler{}

func MakePingTimerHandler(pingResultHandler PingResultHandler, pinger pinger.Pinger, ipaddress string, pingCount int) timer.TimerHandler {
	return &pingHandler{
		resultHandler: pingResultHandler,
		done:          make(chan bool, 1),
		pinger:        pinger,
		ipAddress:     ipaddress,
		pingCount:     pingCount,
	}
}

type pingHandler struct {
	resultHandler PingResultHandler
	done          chan bool
	pinger        pinger.Pinger
	ipAddress     string
	pingCount     int
}

func (p *pingHandler) OnTimeout() {
	result, err := p.pinger.Ping(p.ipAddress, p.pingCount) // block here
	if err != nil {
		logger.Error(err)
		p.Cancel()
		return
	}
	p.resultHandler.OnPingResultHandle(result)
}

func (p *pingHandler) Done() <-chan bool {
	return p.done
}

func (p *pingHandler) Cancel() {
	p.done <- true
}
