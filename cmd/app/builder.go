package app

import (
	"networkmonitor/cmd/config"
	db "networkmonitor/db/kv"
	pingengine "networkmonitor/engine/ping"
	rankengine "networkmonitor/engine/rank"
	"networkmonitor/net/http"
	"networkmonitor/parser"
	"networkmonitor/timer"
)

var _ Builder = &builder{}

type Builder interface {
	BuildServer() http.Server
	BuildPingEngine() pingengine.Engine
	BuildRankEngine() rankengine.Engine
}

type builder struct {
	db     db.KvDb
	config config.Config
}

func MakeBuilder(db db.KvDb, config config.Config) Builder {
	return &builder{
		db:     db,
		config: config,
	}
}

func (b builder) BuildServer() http.Server {
	return http.MakeServer(b.config.ServerIpAddr())
}

func (b builder) BuildPingEngine() pingengine.Engine {
	handler := pingengine.MakePingResultHandler(b.db, parser.MakeJsonParser())
	timerFactory := pingengine.MakePingTimerHandlerFactory()
	return pingengine.MakePingEngine(
		[]pingengine.PingResultHandler{handler},
		timerFactory,
		timer.MakeTimer(),
		b.config.PingEnginePingInterval(),
		b.config.PingEnginePingCount(),
		b.db,
		b.config.PingEngineMaxPingAllowed())
}

func (b builder) BuildRankEngine() rankengine.Engine {
	return rankengine.MakeRankEngine(
		b.config.RankEngineMaxEntryAllowed(),
		b.db,
		timer.MakeTimer(),
		parser.MakeJsonParser())
}
