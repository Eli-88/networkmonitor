package app

import (
	"networkmonitor/cmd/config"
	db "networkmonitor/core/db/kv"
	"networkmonitor/core/net/http"
	"networkmonitor/core/parser"
	"networkmonitor/core/timer"
	"networkmonitor/pingengine"
	"networkmonitor/rankengine"
)

var _ Builder = &builder{}

type Builder interface {
	BuildServer() http.Server
	BuildPingEngine() pingengine.Engine
	BuildRankEngine() rankengine.Engine
}

type builder struct {
	serverAddr string
	db         db.KvDb
	timer      timer.Timer
	config     config.Config
}

func MakeBuilder(serverAddr string, db db.KvDb, timer timer.Timer, config config.Config) Builder {
	return &builder{
		serverAddr: serverAddr,
		db:         db,
		timer:      timer,
		config:     config,
	}
}

func (b builder) BuildServer() http.Server {
	return http.MakeServer(b.serverAddr)
}

func (b builder) BuildPingEngine() pingengine.Engine {
	handler := pingengine.MakePingResultHandler(b.db, parser.MakeJsonParser())
	timerFactory := pingengine.MakePingTimerHandlerFactory()
	return pingengine.MakePingEngine(
		[]pingengine.PingResultHandler{handler},
		timerFactory, b.timer,
		b.config.PingEnginePingInterval(),
		b.config.PingEnginePingCount(),
		b.db,
		b.config.PingEngineMaxPingAllowed())
}

func (b builder) BuildRankEngine() rankengine.Engine {
	return rankengine.MakeRankEngine(
		b.config.RankEngineMaxEntryAllowed(),
		b.db,
		b.timer,
		parser.MakeJsonParser())
}
