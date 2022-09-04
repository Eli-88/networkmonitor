package timer

type IsAlive interface {
	Set(bool)
	Get() bool
}

type TimerHandler interface {
	OnTimeout()
}

type TimerControl interface {
	Cancel()
	IsAlive() bool
}

type Delay int64

type Timer interface {
	DispatchTimerHandler(TimerHandler, Delay) TimerControl
}
