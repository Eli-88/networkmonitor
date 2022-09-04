package config

import "networkmonitor/core/timer"

type Config interface {
	PingEnginePingInterval() timer.Delay
	PingEnginePingCount() int
	PingEngineMaxPingAllowed() int
	RankEngineMaxEntryAllowed() int
	DbPath() string
	ServerIpAddr() string
}
