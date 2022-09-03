package app

import (
	"networkmonitor/cmd/handler"
	"networkmonitor/core/logger"
	"networkmonitor/core/net/transport"
	"networkmonitor/pingengine"
	"networkmonitor/rankengine"
)

type App interface {
	Run()
}

type app struct {
	server         transport.HttpServer
	pingEngine     pingengine.Engine
	rankEngine     rankengine.Engine
	handlerBuilder handler.HandlerBuilder
}

func MakeApp(
	server transport.HttpServer,
	pingEngine pingengine.Engine,
	rankEngine rankengine.Engine,
	handlerBuilder handler.HandlerBuilder,
) App {
	return &app{
		server:         server,
		pingEngine:     pingEngine,
		rankEngine:     rankEngine,
		handlerBuilder: handlerBuilder,
	}
}

func (a *app) Run() {
	a.server.RegisterHttpHandler(
		[]transport.HttpMethod{transport.HTTP_GET, transport.HTTP_POST},
		"/register",
		a.handlerBuilder.BuildRegisterHandler())

	a.server.RegisterHttpHandler(
		[]transport.HttpMethod{transport.HTTP_GET, transport.HTTP_POST},
		"/rank/networkspeed",
		a.handlerBuilder.BuildRankHandler())

	err := a.pingEngine.Run()
	if err != nil {
		logger.Fatal(err)
	}

	a.rankEngine.Run()
	a.server.Run()
}
