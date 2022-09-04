package timer

type IsAlive interface {
	Set(bool)
	Get() bool
}

type TimerHandler interface {
	OnTimeout()
	Done() <-chan bool
	Cancel()
	IsAlive() IsAlive
}

type Delay int64

type Timer interface {
	DispatchTimerHandler(TimerHandler, Delay)
}
