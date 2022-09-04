package handler

import (
	rankengine "networkmonitor/engine/rank"
	"networkmonitor/logger"
	"networkmonitor/net/http"
	"networkmonitor/parser"
)

var _ http.RequestHandler = &rankHandler{}

type rankHandler struct {
	rankEngine rankengine.Engine
	parser     parser.Encoder
}

func MakeRankHandler(rankEngine rankengine.Engine, parser parser.Encoder) http.RequestHandler {
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
