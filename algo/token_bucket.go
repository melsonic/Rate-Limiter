package algo

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity        int
	TokenCount      int
	TokenRefillRate int
	LastRefillTime  time.Time
	mut             sync.Mutex
}

func (t *TokenBucket) RefillBucket() {
  t.mut.Lock()
  defer t.mut.Unlock()
	now := time.Now()
	duration := now.Sub(t.LastRefillTime)
	tokensToAdd := int(duration.Seconds()) * t.TokenRefillRate
	t.TokenCount = min(t.TokenCount+tokensToAdd, t.Capacity)
	t.LastRefillTime = now
}

func (t *TokenBucket) HandleNewRequest() bool {
	t.RefillBucket()
  t.mut.Lock()
  t.mut.Unlock()

	if t.TokenCount == 0 {
		fmt.Printf("count : %d & capacity : %d ... so disallow\n", t.TokenCount, t.Capacity)
		return false
	}
	t.TokenCount = t.TokenCount - 1
	fmt.Printf("count : %d & capacity : %d ... so allow\n", t.TokenCount+1, t.Capacity)
	return true
}
