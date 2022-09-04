package handler

import (
	"networkmonitor/core/net/http"
	"networkmonitor/core/parser"
	"networkmonitor/pingengine"
	"networkmonitor/rankengine"
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
