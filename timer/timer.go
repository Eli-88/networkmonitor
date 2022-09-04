package timer

import "time"

// interface compliance
var _ Timer = &timerImpl{}

type timerImpl struct{}

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
