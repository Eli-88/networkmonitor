package handler

import (
	"networkmonitor/core/net/transport"
	"networkmonitor/core/parser"
	"networkmonitor/pingengine"
	"networkmonitor/rankengine"
)

// interface compliance
var _ HandlerBuilder = handlerBuilder{}

type HandlerBuilder interface {
	BuildRegisterHandler() transport.HttpRequestHandler
	BuildRankHandler() transport.HttpRequestHandler
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

func (h handlerBuilder) BuildRegisterHandler() transport.HttpRequestHandler {
	return MakeRegisterHandler(h.pingEngine, parser.MakeJsonParser())
}

func (h handlerBuilder) BuildRankHandler() transport.HttpRequestHandler {
	return MakeRankHandler(h.rankEngine, parser.MakeJsonParser())
}
