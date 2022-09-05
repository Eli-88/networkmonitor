package app

import (
	"networkmonitor/cmd/handler"
	pingengine "networkmonitor/engine/ping"
	rankengine "networkmonitor/engine/rank"
	"networkmonitor/logger"
	"networkmonitor/net/http"
)

type App interface {
	Run()
}

type app struct {
	server         http.Server
	pingEngine     pingengine.Engine
	rankEngine     rankengine.Engine
	handlerBuilder handler.HandlerBuilder
}

func MakeApp(
	server http.Server,
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
		[]http.Method{http.GET, http.POST},
		"/register",
		a.handlerBuilder.BuildRegisterHandler())

	a.server.RegisterHttpHandler(
		[]http.Method{http.GET, http.POST},
		"/rank/networkspeed",
		a.handlerBuilder.BuildRankHandler())

	err := a.pingEngine.Run()
	if err != nil {
		logger.Fatal(err)
	}

	a.rankEngine.Run()

	err = a.server.Run()
	if err != nil {
		logger.Fatal(err)
	}
}
