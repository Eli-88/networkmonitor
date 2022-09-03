package rankengine

type Engine interface {
	Run()
	TopIpAddrInFastestOrder() RankByTimeReponseCollection
}
