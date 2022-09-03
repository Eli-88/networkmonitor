package app

import (
	db "networkmonitor/core/db/kv"
	"networkmonitor/core/net/transport"
	"networkmonitor/core/parser"
	"networkmonitor/core/timer"
	"networkmonitor/pingengine"
	"networkmonitor/rankengine"
)

var _ Builder = &builder{}

type Builder interface {
	BuildHttpServer() transport.HttpServer
	BuildPingEngine() pingengine.Engine
	BuildRankEngine() rankengine.Engine
}

type builder struct {
	serverAddr string
	db         db.KvDb
	timer      timer.Timer
}

func MakeBuilder(serverAddr string, db db.KvDb, timer timer.Timer) Builder {
	return &builder{
		serverAddr: serverAddr,
		db:         db,
		timer:      timer,
	}
}

func (b builder) BuildHttpServer() transport.HttpServer {
	return transport.MakeHttpServer(b.serverAddr)
}

func (b builder) BuildPingEngine() pingengine.Engine {
	handler := pingengine.MakePingResultHandler(b.db, parser.MakeJsonParser())
	timerFactory := pingengine.MakePingTimerHandlerFactory()
	return pingengine.MakePingEngine([]pingengine.PingResultHandler{handler}, timerFactory, b.timer, 5000, 10, b.db, 1000)
}

func (b builder) BuildRankEngine() rankengine.Engine {
	return rankengine.MakeRankEngine(1000, b.db, b.timer, parser.MakeJsonParser())
}
