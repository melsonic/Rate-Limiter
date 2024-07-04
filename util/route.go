package util

import (
	"net/http"

	"github.com/melsonic/rate-limiter/util/middleware"
)

func Route(option string) *http.ServeMux {
	router := http.NewServeMux()
	switch option {
	case "1":
		router.HandleFunc("/limited", middleware.TokenBucketMiddlewareRL(LimitedHandler))
	case "2":
		router.HandleFunc("/limited", middleware.FixedWindowCounterMiddlewareRL(LimitedHandler))
		middleware.FixedWindowCounterHelper()
	case "3":
		router.HandleFunc("/limited", middleware.SlidingWindowLogMiddlewareRL(LimitedHandler))
	case "4":
		router.HandleFunc("/limited", middleware.SlidingWindowCounterMiddlewareRL(LimitedHandler))
		middleware.SlidingWindowCounterHelper()
	default:
		router.HandleFunc("/limited", LimitedHandler)
	}
	router.HandleFunc("/unlimited", UnlimitedHandler)
	return router
}
