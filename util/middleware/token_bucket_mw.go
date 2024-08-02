package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/melsonic/rate-limiter/algo"
	"github.com/melsonic/rate-limiter/constants"
)

var BucketList map[string]*algo.TokenBucket = make(map[string]*algo.TokenBucket)

func TokenBucketMiddlewareRL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClientIP := strings.Split(r.Header.Get("X-FORWARDED-FOR"), ", ")[0]
		if ClientIP == "" {
			ClientIP = r.RemoteAddr
		}
		constants.Mut.RLock()
		bucket, bucketPresent := BucketList[ClientIP]
		constants.Mut.RUnlock()
		if !bucketPresent {
			constants.Mut.Lock()
			BucketList[ClientIP] = &algo.TokenBucket{
				Capacity:        constants.BucketCapacity,
				TokenCount:      1,
				TokenRefillRate: constants.TokenRefillRate,
				LastRefillTime:  time.Now(),
			}
			bucket = BucketList[ClientIP]
			constants.Mut.Unlock()
		}
		fmt.Printf("%s   ", ClientIP)
		allowRequest := bucket.HandleIncomingRequest()
		if !allowRequest {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, "Too many requests\n")
			return
		}
		next(w, r)
	}
}
