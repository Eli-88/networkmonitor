package handler

import (
	"networkmonitor/core/logger"
	"networkmonitor/core/net/transport"
	"networkmonitor/core/parser"
	"networkmonitor/rankengine"
)

var _ transport.HttpRequestHandler = &rankHandler{}

type rankHandler struct {
	rankEngine rankengine.Engine
	parser     parser.Encoder
}

func MakeRankHandler(rankEngine rankengine.Engine, parser parser.Encoder) transport.HttpRequestHandler {
	return &rankHandler{
		rankEngine: rankEngine,
		parser:     parser,
	}
}

func (r *rankHandler) OnHttpRequest(body []byte) (string, error) {
	result := r.rankEngine.TopIpAddrInFastestOrder()
	logger.Debug(result)
	response, err := r.parser.Marshal(result)
	if err != nil {
		logger.Error(err)
		return "", nil
	}
	return string(response), nil
}
