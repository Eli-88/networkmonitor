package pingengine

import (
	db "networkmonitor/db/kv"
	"networkmonitor/logger"
	"networkmonitor/net/pinger"
	"networkmonitor/timer"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func MakePingEngine(
	pingResultHandlers []PingResultHandler,
	pingHandlerFactory PingTimerHandlerFactory,
	pingTimer timer.Timer,
	pingInterval timer.Delay,
	pingCount int,
	db db.KvDb,
	maxPingAllowed int,
) Engine {
	return &pingEngine{
		pingResultHandlers:  pingResultHandlers,
		pingHanderFactory:   pingHandlerFactory,
		pingTimer:           pingTimer,
		pingInterval:        pingInterval,
		pingCount:           pingCount,
		pingTimerHandlerMap: cmap.New[timer.TimerHandler](),
		db:                  db,
		maxPingAllowed:      maxPingAllowed,
	}

}

type pingEngine struct {
	pingResultHandlers  []PingResultHandler
	pingHanderFactory   PingTimerHandlerFactory
	pingTimer           timer.Timer
	pingInterval        timer.Delay
	pingCount           int
	pingTimerHandlerMap cmap.ConcurrentMap[timer.TimerHandler]
	db                  db.KvDb
	maxPingAllowed      int
}

func (p *pingEngine) RegisterIpAddress(ipAddress string) bool {
	logger.Debug("registering address:", ipAddress, "at ping engine addr", &p)
	_, exist := p.pingTimerHandlerMap.Get(ipAddress)
	if !exist {
		pingHandler := p.pingHanderFactory.CreatePingTimerHandler(p, ipAddress, p.pingCount)
		p.pingTimer.DispatchTimerHandler(pingHandler, p.pingInterval)
		p.pingTimerHandlerMap.Set(ipAddress, pingHandler)

		return true
	} else {
		return false
	}
}

func (p *pingEngine) OnPingResultHandle(result pinger.Stats) {
	logger.Trace()
	for _, handler := range p.pingResultHandlers {
		handler.OnPingResultHandle(result)
	}
}

func (p *pingEngine) Run() error {
	logger.Trace()
	err := p.db.GetAllKeyValue(func(key []byte, value []byte) bool {
		if len(p.pingTimerHandlerMap) < p.maxPingAllowed {
			p.RegisterIpAddress(string(key))
			return true
		} else {
			return false
		}
	})

	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
