package algo

import (
	"fmt"
	"time"

	"github.com/melsonic/rate-limiter/constants"
)

type SlidingWindowCounterEntry struct {
	CurrentWindowCount int
	PrevWindowCount    int
	WindowStartTime    time.Time
	WindowEndTime      time.Time
}

func (e *SlidingWindowCounterEntry) WindowReset() {
	e.PrevWindowCount = e.CurrentWindowCount
	e.CurrentWindowCount = 0
	e.WindowStartTime = e.WindowEndTime
	e.WindowEndTime = e.WindowStartTime.Add(time.Second * constants.SlidingWindowCounter_WindowSize)
}

func (e *SlidingWindowCounterEntry) HandleIncomingRequest() bool {
	var now time.Time = time.Now()
	var durationFromWindowStart time.Duration = now.Sub(e.WindowStartTime)
	var percentThroughWindow float32 = float32(durationFromWindowStart.Seconds()) / (constants.SlidingWindowCounter_WindowSize)
	var windowCount float32 = percentThroughWindow*float32(e.CurrentWindowCount) + (1-percentThroughWindow)*float32(e.PrevWindowCount)
	if windowCount >= constants.SlidingWindowCounter_WindowThreshold {
		fmt.Printf("%f   %v   ... so disallow\n", windowCount, constants.SlidingWindowCounter_WindowThreshold)
		return false
	}
	e.CurrentWindowCount = e.CurrentWindowCount + 1
	fmt.Printf("%f   %v   ... so allow\n", windowCount, constants.SlidingWindowCounter_WindowSize)
	return true
}
