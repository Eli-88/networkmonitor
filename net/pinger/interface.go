package pinger

type Stats interface {
	// PacketsRecv is the number of packets received.
	PacketsRecv() int

	// PacketsSent is the number of packets sent.
	PacketsSent() int

	// PacketsRecvDuplicates is the number of duplicate responses there were to a sent packet.
	PacketsRecvDuplicates() int

	// PacketLoss is the percentage of packets lost.
	PacketLoss() float64

	// Addr is the string address of the host being pinged.
	Addr() string

	// MinRtt is the minimum round-trip time sent via this pinger.
	MinRtt() int64

	// MaxRtt is the maximum round-trip time sent via this pinger.
	MaxRtt() int64

	// AvgRtt is the average round-trip time sent via this pinger.
	AvgRtt() int64

	// StdDevRtt is the standard deviation of the round-trip times sent via
	// this pinger.
	StdDevRtt() int64
}

type Pinger interface {
	Ping(string, int) (Stats, error)
}
