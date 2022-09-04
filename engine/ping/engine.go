package pingengine

import (
	db "networkmonitor/db/kv"
	"networkmonitor/logger"
	"networkmonitor/net/pinger"
	"networkmonitor/parser"
	"networkmonitor/timer"
)

func MakePingEngine(
	pingTimerHandler timer.TimerHandler,
	pingTimer timer.Timer,
	pingInterval timer.Delay,
	pingCount int,
	db db.KvDb,
	maxPingAllowed int,
	parser parser.Encoder,
) Engine {
	return &pingEngine{
		pingTimerHandler: pingTimerHandler,
		pingTimer:        pingTimer,
		pingInterval:     pingInterval,
		pingCount:        pingCount,
		db:               db,
		maxPingAllowed:   maxPingAllowed,
		parser:           parser,
	}

}

type pingEngine struct {
	pingTimerHandler timer.TimerHandler
	pingTimer        timer.Timer
	pingInterval     timer.Delay
	pingCount        int
	db               db.KvDb
	maxPingAllowed   int
	parser           parser.Encoder
}

func (p *pingEngine) RegisterIpAddress(ipAddress string) {
	logger.Debug("registering address:", ipAddress)

	go func() {
		// need to be successful before registering
		stats, err := pinger.MakePinger().Ping(ipAddress, p.pingCount)
		if err != nil {
			logger.Error(err)
			return
		}
		response, err := p.parser.Marshal(stats)
		if err != nil {
			logger.Error(err)
			return
		}
		p.db.SetKvKeyValue([]byte(ipAddress), response)
	}()
}

func (p *pingEngine) Run() error {
	logger.Trace()
	p.pingTimer.DispatchTimerHandler(p.pingTimerHandler, p.pingInterval)
	return nil
}
