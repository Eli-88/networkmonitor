package rankengine

import (
	"container/heap"
	"errors"
	"networkmonitor/logger"
	"sort"
)

// interface compliance
var _ heap.Interface = &RankByTimeReponseCollection{}
var _ sort.Interface = &RankByTimeReponseCollection{}

type RankByTimeReponse struct {
	Addr   string `json:"IpAddress`
	AvgRtt int64  `json:"AverageRtt"`
}

type RankByTimeReponseCollection []RankByTimeReponse

func (r RankByTimeReponseCollection) Len() int {
	return len(r)
}

func (r RankByTimeReponseCollection) Less(i, j int) bool {
	return r[i].AvgRtt < r[j].AvgRtt
}

func (r RankByTimeReponseCollection) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r *RankByTimeReponseCollection) Push(x any) {
	item := x.(RankByTimeReponse)
	*r = append(*r, item)
}

func (r *RankByTimeReponseCollection) Pop() any {
	old := *r
	n := len(old)
	if n == 0 {
		err := errors.New("need to check len before pop")
		logger.Fatal(err)
	}
	item := old[n-1]
	*r = old[0 : n-1]
	return item
}
