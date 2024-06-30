package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/melsonic/rate-limiter/algo"
	"github.com/melsonic/rate-limiter/constants"
)

var (
	BucketList map[string]*algo.TokenBucket = make(map[string]*algo.TokenBucket)
	mut        sync.RWMutex
)

func TokenBucketMiddlewareRL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClientIP := r.RemoteAddr
		mut.RLock()
		bucket, bucketPresent := BucketList[ClientIP]
		mut.RUnlock()
		if !bucketPresent {
      mut.Lock()
			BucketList[ClientIP] = &algo.TokenBucket{
				Capacity:        constants.BucketCapacity,
				TokenCount:      1,
				TokenRefillRate: constants.TokenRefillRate,
				LastRefillTime:  time.Now(),
			}
			bucket = BucketList[ClientIP]
      mut.Unlock()
		}
		fmt.Printf("%s   ", ClientIP)
		allowRequest := bucket.HandleNewRequest()
		if !allowRequest {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, "Too many requests\n")
			return
		}
		next(w, r)
	}
}
