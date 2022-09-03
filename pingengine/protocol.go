package pingengine

type PingStats struct {
	// PacketsRecv is the number of packets received.
	PacketsRecv int `json:"PacketsRecv"`

	// PacketsSent is the number of packets sent.
	PacketsSent int `json:"PacketsSent"`

	// PacketsRecvDuplicates is the number of duplicate responses there were to a sent packet.
	PacketsRecvDuplicates int `json:"PacketsRecvDuplicates"`

	// PacketLoss is the percentage of packets lost.
	PacketLoss float64 `json:"PacketLoss"`

	// Addr is the string address of the host being pinged.
	Addr string `json:"Addr"`

	// MinRtt is the minimum round-trip time sent via this pinger.
	MinRtt int64 `json:"MinRtt"`

	// MaxRtt is the maximum round-trip time sent via this pinger.
	MaxRtt int64 `json:"MaxRtt"`

	// AvgRtt is the average round-trip time sent via this pinger.
	AvgRtt int64 `json:"AvgRtt"`

	// StdDevRtt is the standard deviation of the round-trip times sent via
	// this pinger.
	StdDevRtt int64 `json:"StdDevRtt"`
}
