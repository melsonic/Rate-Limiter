package algo

import (
	"fmt"

	"github.com/melsonic/rate-limiter/constants"
)

type FixedWindowEntry struct {
	CurrentRequestCount int
}

func (e *FixedWindowEntry) Reset() {
	e.CurrentRequestCount = 0
}

func (e *FixedWindowEntry) IncrementCounter() {
	e.CurrentRequestCount = e.CurrentRequestCount + 1
}

func (e *FixedWindowEntry) HandleIncomingRequest() bool {
	if e.CurrentRequestCount >= constants.FixedWindowCounter_RequestThreshold {
		fmt.Printf("%d   ...   so disallow\n", e.CurrentRequestCount)
		return false
	}
	e.IncrementCounter()
	fmt.Printf("%d   ...   so allow\n", e.CurrentRequestCount)
	return true
}
