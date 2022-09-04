package config

import "networkmonitor/timer"

type pingEngineConfig struct {
	PingInterval   timer.Delay `json:"PingInterval"`
	PingCount      int         `json:"PingCount"`
	MaxPingAllowed int         `json:"MaxPingAllowed"`
}
