package handler

import (
	pingengine "networkmonitor/engine/ping"
	rankengine "networkmonitor/engine/rank"
	"networkmonitor/net/http"
	"networkmonitor/parser"
)

// interface compliance
var _ HandlerBuilder = handlerBuilder{}

type HandlerBuilder interface {
	BuildRegisterHandler() http.RequestHandler
	BuildRankHandler() http.RequestHandler
}

func MakeHandlerBuilder(pingEngine pingengine.Engine, rankEngine rankengine.Engine) HandlerBuilder {
	return &handlerBuilder{
		pingEngine: pingEngine,
		rankEngine: rankEngine,
	}
}

type handlerBuilder struct {
	pingEngine pingengine.Engine
	rankEngine rankengine.Engine
}

func (h handlerBuilder) BuildRegisterHandler() http.RequestHandler {
	return MakeRegisterHandler(h.pingEngine, parser.MakeJsonParser())
}

func (h handlerBuilder) BuildRankHandler() http.RequestHandler {
	return MakeRankHandler(h.rankEngine, parser.MakeJsonParser())
}
