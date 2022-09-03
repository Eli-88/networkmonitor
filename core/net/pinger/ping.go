package pinger

import (
	"github.com/go-ping/ping"
)

// interface compliance
var _ Pinger = pingerImpl{}
var _ Stats = statistics{}

type pingerImpl struct{}

type statistics struct {
	// PacketsRecv is the number of packets received.
	packetsRecv int

	// PacketsSent is the number of packets sent.
	packetsSent int

	// PacketsRecvDuplicates is the number of duplicate responses there were to a sent packet.
	packetsRecvDuplicates int

	// PacketLoss is the percentage of packets lost.
	packetLoss float64

	// Addr is the string address of the host being pinged.
	addr string

	// MinRtt is the minimum round-trip time sent via this pinger.
	minRtt int64

	// MaxRtt is the maximum round-trip time sent via this pinger.
	maxRtt int64

	// AvgRtt is the average round-trip time sent via this pinger.
	avgRtt int64

	// StdDevRtt is the standard deviation of the round-trip times sent via
	// this pinger.
	stdDevRtt int64
}

func MakePinger() Pinger {
	return &pingerImpl{}
}

func (p pingerImpl) Ping(ipAddress string, count int) (Stats, error) {
	pinger, err := ping.NewPinger(ipAddress)
	if err != nil {
		return nil, MakePingInitError(err)
	}
	pinger.Count = count
	err = pinger.Run() // block here
	if err != nil {
		return nil, MakePingError(err)
	}
	stats := pinger.Statistics()
	return &statistics{
		packetsRecv:           stats.PacketsRecv,
		packetsSent:           stats.PacketsSent,
		packetLoss:            stats.PacketLoss,
		packetsRecvDuplicates: stats.PacketsRecvDuplicates,
		addr:                  stats.Addr,
		minRtt:                stats.MinRtt.Milliseconds(),
		maxRtt:                stats.MaxRtt.Milliseconds(),
		avgRtt:                stats.AvgRtt.Milliseconds(),
		stdDevRtt:             stats.StdDevRtt.Milliseconds(),
	}, nil
}

type pingStats struct {
	ipAddress            string
	packetSentCount      int
	packetRecvCount      int
	packetLossPercentage float64
	averageRoundTripTime int64
}

// PacketsRecv is the number of packets received.
func (s statistics) PacketsRecv() int {
	return s.packetsRecv
}

// PacketsSent is the number of packets sent.
func (s statistics) PacketsSent() int {
	return s.packetsSent
}

// PacketsRecvDuplicates is the number of duplicate responses there were to a sent packet.
func (s statistics) PacketsRecvDuplicates() int {
	return s.packetsRecvDuplicates
}

// PacketLoss is the percentage of packets lost.
func (s statistics) PacketLoss() float64 {
	return s.packetLoss
}

// Addr is the string address of the host being pinged.
func (s statistics) Addr() string {
	return s.addr
}

// MinRtt is the minimum round-trip time sent via this pinger.
func (s statistics) MinRtt() int64 {
	return s.minRtt
}

// MaxRtt is the maximum round-trip time sent via this pinger.
func (s statistics) MaxRtt() int64 {
	return s.maxRtt
}

// AvgRtt is the average round-trip time sent via this pinger.
func (s statistics) AvgRtt() int64 {
	return s.avgRtt
}

// StdDevRtt is the standard deviation of the round-trip times sent via
// this pinger.
func (s statistics) StdDevRtt() int64 {
	return s.stdDevRtt
}
