package algo

import (
	"fmt"
	"time"

	"github.com/melsonic/rate-limiter/constants"
)

type SlidingWindowLog struct {
	TimeStamps []time.Time
}

func (s *SlidingWindowLog) AddTimeStamp(t time.Time) {
	s.TimeStamps = append(s.TimeStamps, t)
}

func (s *SlidingWindowLog) RemoveRedundantTimeStamps(t time.Time) {
	startTimestamp := t.Add(time.Second * constants.SlidingWindowLog_WindowSize * (-1))
	var result []time.Time
	for _, timestamp := range s.TimeStamps {
		if timestamp.After(startTimestamp) {
			result = append(result, timestamp)
		}
	}
	copy(s.TimeStamps, result)
}

func (s *SlidingWindowLog) HandleIncomingRequest() bool {
	now := time.Now()
	s.RemoveRedundantTimeStamps(now)
	s.AddTimeStamp(now)

	if len(s.TimeStamps) >= constants.SlidingWindowLog_WindowThreshold {
		fmt.Printf("%v   %v    ...   so disallow\n", len(s.TimeStamps), constants.SlidingWindowLog_WindowThreshold)
		return false
	}
	fmt.Printf("%v   %v    ...   so allow\n", len(s.TimeStamps), constants.SlidingWindowLog_WindowThreshold)
	return true
}
