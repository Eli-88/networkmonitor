package handler

import (
	"networkmonitor/core/logger"
	"networkmonitor/core/net/transport"
	"networkmonitor/core/parser"
	"networkmonitor/pingengine"
)

var _ transport.HttpRequestHandler = registerHandler{}

type registerProtocol struct {
	Addr string `json:"IpAddress"`
}

type registerHandler struct {
	pingEngine pingengine.Engine
	parser     parser.Parser
}

func MakeRegisterHandler(pingEngine pingengine.Engine, parser parser.Parser) transport.HttpRequestHandler {
	return &registerHandler{
		pingEngine: pingEngine,
		parser:     parser,
	}
}

func (r registerHandler) OnHttpRequest(body []byte) (string, error) {
	logger.Debug("recv:", string(body))
	msg := &registerProtocol{}
	err := r.parser.Unmarshal([]byte(body), msg)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	if r.pingEngine.RegisterIpAddress(msg.Addr) {
		return "Ok", nil
	} else {
		return "Already Exists", nil
	}
}
