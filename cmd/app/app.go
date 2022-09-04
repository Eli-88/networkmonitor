package app

import (
	"networkmonitor/cmd/handler"
	"networkmonitor/core/logger"
	"networkmonitor/core/net/http"
	"networkmonitor/pingengine"
	"networkmonitor/rankengine"
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
	a.server.Run()
}
