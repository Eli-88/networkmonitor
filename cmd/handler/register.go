package handler

import (
	pingengine "networkmonitor/engine/ping"
	"networkmonitor/logger"
	"networkmonitor/net/http"
	"networkmonitor/parser"
)

var _ http.RequestHandler = registerHandler{}

type registerProtocol struct {
	Addr string `json:"IpAddress"`
}

type registerHandler struct {
	pingEngine pingengine.Engine
	parser     parser.Parser
}

func MakeRegisterHandler(pingEngine pingengine.Engine, parser parser.Parser) http.RequestHandler {
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

	r.pingEngine.RegisterIpAddress(msg.Addr)
	return "Ok", nil

}
