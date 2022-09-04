package config

import (
	"fmt"
	"io/ioutil"
	"networkmonitor/parser"
	"networkmonitor/timer"
)

var _ Config = &allConfig{}

func MakeConfig(filename string) (Config, error) {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, makeConfigReadError(err)
	}

	cfg := &allConfig{}
	decoder := parser.MakeJsonParser()
	err = decoder.Unmarshal(content, cfg)

	if err != nil {
		return nil, makeConfigParseError(err)
	}

	return cfg, nil
}

type allConfig struct {
	Db               dbConfig         `json:"Db"`
	PingEngineConfig pingEngineConfig `json:"PingEngine"`
	RankEngineConfig rankEngineConfig `json:"RankEngine"`
	ServerConfig     serverConfig     `json:"Server"`
}

func (a allConfig) PingEnginePingInterval() timer.Delay {
	return a.PingEngineConfig.PingInterval
}

func (a allConfig) PingEnginePingCount() int {
	return a.PingEngineConfig.PingCount
}

func (a allConfig) PingEngineMaxPingAllowed() int {
	return a.PingEngineConfig.MaxPingAllowed
}

func (a allConfig) RankEngineMaxEntryAllowed() int {
	return a.RankEngineConfig.MaxEntryAllowed
}

func (a allConfig) DbPath() string {
	return a.Db.Path
}

func (a allConfig) ServerIpAddr() string {
	return fmt.Sprintf("%s:%d", a.ServerConfig.Host, a.ServerConfig.Port)
}
