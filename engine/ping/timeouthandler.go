package pingengine

import (
	db "networkmonitor/db/kv"
	"networkmonitor/logger"
	"networkmonitor/net/pinger"
	"networkmonitor/timer"
)

// interface compliance
var _ timer.TimerHandler = &pingHandler{}

func MakePingTimerHandler(pingResultHandler PingResultHandler, pinger pinger.Pinger, db db.KvDb, pingCount int) timer.TimerHandler {
	return &pingHandler{
		resultHandler: pingResultHandler,
		done:          make(chan bool, 1),
		pinger:        pinger,
		db:            db,
		pingCount:     pingCount,
	}
}

type pingHandler struct {
	resultHandler PingResultHandler
	done          chan bool
	pinger        pinger.Pinger
	db            db.KvDb
	pingCount     int
}

func (p *pingHandler) OnTimeout() {
	p.db.GetAllKeyValue(func(ipaddress, value []byte) bool {
		go func() {
			result, err := p.pinger.Ping(string(ipaddress), p.pingCount) // block here
			if err != nil {
				logger.Error(err)
				return
			}
			p.resultHandler.OnPingResultHandle(result)
		}()
		return true
	})
}
