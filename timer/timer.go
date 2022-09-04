package timer

import (
	"sync/atomic"
	"time"
)

// interface compliance
var _ Timer = &timerImpl{}

type timerImpl struct{}

type isAlive struct {
	flag atomic.Value
}

func MakeTimer() Timer {
	return &timerImpl{}
}

func (t timerImpl) DispatchTimerHandler(handler TimerHandler, delayInMilliseconds Delay) {
	ticker := time.NewTicker(time.Duration(delayInMilliseconds) * time.Millisecond)

	go func() {
		for {
			select {
			case <-handler.Done():
				return
			case <-ticker.C:
				handler.OnTimeout()
			}
		}
	}()
}

func (i *isAlive) Set(flag bool) {
	i.flag.Store(flag)
}

func (i *isAlive) Get() bool {
	return i.flag.Load().(bool)
}

func MakeIsAlive(flag bool) IsAlive {
	a := &isAlive{}
	a.flag.Store(flag)
	return a
}
