package util

import "net/http"
import "github.com/melsonic/rate-limiter/util/middleware"

func Route() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/limited", middleware.TokenBucketMiddlewareRL(LimitedHandler))
	router.HandleFunc("/unlimited", UnlimitedHandler)
	return router
}
