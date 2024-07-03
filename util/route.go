package util

import (
	"net/http"
	"time"

	"github.com/melsonic/rate-limiter/constants"
	"github.com/melsonic/rate-limiter/util/middleware"
)

func FixedWindowHelper() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case t := <-ticker.C:
				if t == constants.FixedWindowCounter_EndTimeStamp {
					constants.Mut.Lock()
					for _, v := range middleware.FixedWindowCounterList {
						v.Reset()
					}
					constants.Mut.Unlock()
					constants.FixedWindowCounter_EndTimeStamp = constants.FixedWindowCounter_EndTimeStamp.Add(time.Second * constants.FixedWindowCounter_WindowSize)
				}
			}
		}
	}()
}

func Route(option string) *http.ServeMux {
	router := http.NewServeMux()
	switch option {
    case "1":
      router.HandleFunc("/limited", middleware.TokenBucketMiddlewareRL(LimitedHandler))
    case "2":
      router.HandleFunc("/limited", middleware.FixedWindowCounterMiddlewareRL(LimitedHandler))
      FixedWindowHelper()
    case "3":
      router.HandleFunc("/limited", middleware.SlidingWindowLogMiddlewareRL(LimitedHandler))
    default:
      router.HandleFunc("/limited", LimitedHandler)
	}
	router.HandleFunc("/unlimited", UnlimitedHandler)
	return router
}
