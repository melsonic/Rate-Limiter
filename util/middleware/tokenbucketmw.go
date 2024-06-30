package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/melsonic/rate-limiter/algo"
	"github.com/melsonic/rate-limiter/constants"
)

var BucketList map[string]algo.TokenBucket = make(map[string]algo.TokenBucket)

func TokenBucketMiddlewareRL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClientIP := r.RemoteAddr
		bucket, bucketPresent := BucketList[ClientIP]
		if !bucketPresent {
			BucketList[ClientIP] = algo.TokenBucket{Capacity: constants.BucketCapacity, TokenCount: 0, TokenRefillRate: constants.TokenRefillRate, LastRefillTime: time.Now()}
		}
		bucket = BucketList[ClientIP]
		fmt.Printf("%s => ", ClientIP)
		if !bucket.HandleNewRequest() {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, "Too many requests\n")
			return
		}
		next(w, r)
	}
}
