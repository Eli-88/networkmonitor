package timer

import (
	"networkmonitor/logger"
	"sync/atomic"
	"time"
)

// interface compliance
var _ Timer = &timerImpl{}
var _ TimerControl = &timerControl{}

type timerImpl struct{}

type timerControl struct {
	handler TimerHandler
	done    chan bool
	isAlive atomic.Value
}

func MakeTimer() Timer {
	return &timerImpl{}
}

func makeTimerControl(handler TimerHandler) *timerControl {
	t := &timerControl{
		handler: handler,
		done:    make(chan bool, 1),
	}
	t.isAlive.Store(true)
	return t
}

func (t timerImpl) DispatchTimerHandler(handler TimerHandler, delayInMilliseconds Delay) TimerControl {
	ticker := time.NewTicker(time.Duration(delayInMilliseconds) * time.Millisecond)

	control := makeTimerControl(handler)
	go func() {
		for {
			select {
			case <-control.Done():
				return
			case <-ticker.C:
				control.OnTimeout()
			}
		}
	}()
	return control
}

func (t timerControl) Done() <-chan bool {
	return t.done
}

func (t *timerControl) Cancel() {
	t.done <- true
	t.isAlive.Store(false)
}

func (t timerControl) IsAlive() bool {
	return t.isAlive.Load().(bool)
}

func (t timerControl) OnTimeout() {
	logger.Debug("timeout triggered on handler[", &t.handler, "]")
	t.handler.OnTimeout()
}
