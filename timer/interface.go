package timer

type TimerHandler interface {
	OnTimeout()
	Done() <-chan bool
	Cancel()
}

type Delay int64

type Timer interface {
	DispatchTimerHandler(TimerHandler, Delay)
}
