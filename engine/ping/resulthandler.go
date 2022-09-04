package pingengine

import (
	db "networkmonitor/db/kv"
	"networkmonitor/logger"
	"networkmonitor/net/pinger"
	"networkmonitor/parser"
)

var _ PingResultHandler = &resultHandler{}

func MakePingResultHandler(kv db.KvDb, parser parser.Encoder) PingResultHandler {
	return &resultHandler{kv: kv, parser: parser}
}

type resultHandler struct {
	kv     db.KvDb
	parser parser.Encoder
}

func (r *resultHandler) OnPingResultHandle(result pinger.Stats) {
	logger.Debug("ping result:", result)
	key := result.Addr()

	response, err := r.parser.Marshal(PingStats{
		PacketsRecv:           result.PacketsRecv(),
		PacketsSent:           result.PacketsSent(),
		PacketsRecvDuplicates: result.PacketsRecvDuplicates(),
		PacketLoss:            result.PacketLoss(),
		Addr:                  result.Addr(),
		MinRtt:                result.MinRtt(),
		AvgRtt:                result.AvgRtt(),
		MaxRtt:                result.MaxRtt(),
		StdDevRtt:             result.StdDevRtt(),
	})

	if err != nil {
		goto FAIL
	}

	err = r.kv.UpdateKvExistingValue([]byte(key), response)
	if err != nil {
		goto FAIL
	}

	return
FAIL:
	logger.Error(err)

}
