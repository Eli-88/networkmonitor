package rankengine

import (
	"container/heap"
	db "networkmonitor/db/kv"
	pingengine "networkmonitor/engine/ping"
	"networkmonitor/logger"
	"networkmonitor/parser"
	"networkmonitor/timer"
	"sort"
	"sync/atomic"
)

//interface compliance
var _ Engine = &rankEngine{}
var _ timer.TimerHandler = &rankEngine{}

type rankEngine struct {
	maxEntryAllowed int
	results         atomic.Value
	db              db.KvDb
	timer           timer.Timer
	done            chan bool
	parser          parser.Decoder
	isAlive         bool
}

func MakeRankEngine(maxEntryAllowed int, db db.KvDb, timer timer.Timer, parser parser.Decoder) Engine {
	retVal := &rankEngine{
		maxEntryAllowed: maxEntryAllowed,
		results:         atomic.Value{},
		db:              db,
		timer:           timer,
		done:            make(chan bool, 1),
		parser:          parser,
		isAlive:         true,
	}
	retVal.results.Store(RankByTimeReponseCollection{})
	return retVal
}

func (r *rankEngine) TopIpAddrInFastestOrder() RankByTimeReponseCollection {
	logger.Debug(r.results.Load().(RankByTimeReponseCollection))
	return r.results.Load().(RankByTimeReponseCollection)
}

func (r *rankEngine) Run() {
	logger.Trace()
	r.timer.DispatchTimerHandler(r, timer.Delay(5000))
}

func (r *rankEngine) OnTimeout() {
	result := RankByTimeReponseCollection{}
	err := r.db.GetAllKeyValue(func(key, value []byte) bool {
		stats := pingengine.PingStats{}
		r.parser.Unmarshal(value, &stats)

		var timePenalty int64 = (int64)(stats.PacketLoss * float64(stats.AvgRtt))
		avgRtt := timePenalty + stats.AvgRtt

		heap.Push(&result, RankByTimeReponse{Addr: string(key), AvgRtt: avgRtt})
		if len(result) > r.maxEntryAllowed {
			heap.Pop(&result)
		}

		return true
	})

	if err != nil {
		logger.Error(err)
	}
	logger.Debug("before stored result:", r.results.Load().(RankByTimeReponseCollection))
	sort.Sort(result)
	r.results.Store(result)
	logger.Debug("stored result:", r.results.Load().(RankByTimeReponseCollection))
}

func (r *rankEngine) Done() <-chan bool {
	return r.done
}

func (r *rankEngine) Cancel() {
	r.done <- true
	r.isAlive = false
}

func (r rankEngine) IsAlive() bool {
	return r.isAlive
}
