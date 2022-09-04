package app

import (
	"networkmonitor/cmd/config"
	db "networkmonitor/db/kv"
	pingengine "networkmonitor/engine/ping"
	rankengine "networkmonitor/engine/rank"
	"networkmonitor/net/http"
	"networkmonitor/net/pinger"
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
	pingTimerHandler := pingengine.MakePingTimerHandler(
		pingengine.MakePingResultHandler(
			b.db,
			parser.MakeJsonParser()),
		pinger.MakePinger(),
		b.db,
		b.config.PingEnginePingCount(),
	)
	return pingengine.MakePingEngine(
		pingTimerHandler,
		timer.MakeTimer(),
		b.config.PingEnginePingInterval(),
		b.config.PingEnginePingCount(),
		b.db,
		b.config.PingEngineMaxPingAllowed(),
		parser.MakeJsonParser())
}

func (b builder) BuildRankEngine() rankengine.Engine {
	return rankengine.MakeRankEngine(
		b.config.RankEngineMaxEntryAllowed(),
		b.db,
		timer.MakeTimer(),
		parser.MakeJsonParser())
}
