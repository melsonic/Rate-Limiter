package constants

import (
	"fmt"
	"sync"
	"time"
)

const PORT int = 3045

// token bucket algorithm
const (
	BucketCapacity  int = 10
	TokenRefillRate int = 1
)

// fixed window counter algorithm
const (
	FixedWindowCounter_WindowSize       = 10
	FixedWindowCounter_RequestThreshold = 10
)

var FixedWindowCounter_EndTimeStamp time.Time

// sliding window log algorithm
const (
	SlidingWindowLog_WindowSize      = 10
	SlidingWindowLog_WindowThreshold = 10
)

// sliding window counter algorithm
const (
	SlidingWindowCounter_WindowSize      = 4
	SlidingWindowCounter_WindowThreshold = 3
)

var SlidingWindowCounter_WindowStartTime time.Time

// global values
var (
	ServerAddr string = fmt.Sprintf(":%d", PORT)
	Mut        sync.RWMutex
)

var AlgorithmOptionsArray = []string{"Enter 1, 2 as per choice", "1) Token Bucket", "2) Fixed Window", "3) Sliding Window Log", "4) Sliding Window Counter"}
