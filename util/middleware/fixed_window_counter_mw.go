package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/melsonic/rate-limiter/algo"
	"github.com/melsonic/rate-limiter/constants"
)

var FixedWindowCounterList map[string]*algo.FixedWindowEntry = make(map[string]*algo.FixedWindowEntry)

func FixedWindowCounterMiddlewareRL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClientIP := strings.Split(r.Header.Get("X-FORWARDED-FOR"), ", ")[0]
		if ClientIP == "" {
			ClientIP = r.RemoteAddr
		}
		constants.Mut.RLock()
		entry, entryPresent := FixedWindowCounterList[ClientIP]
		constants.Mut.RUnlock()
		if !entryPresent {
			entry = &algo.FixedWindowEntry{CurrentRequestCount: 1}
			constants.Mut.Lock()
			FixedWindowCounterList[ClientIP] = entry
			constants.Mut.Unlock()
		}
		fmt.Printf("%s   ", ClientIP)
		var allowRequest bool = entry.HandleIncomingRequest()
		if !allowRequest {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, "Too many requests\n")
			return
		}
		next(w, r)
	}
}

func FixedWindowCounterHelper() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case t := <-ticker.C:
				if t == constants.FixedWindowCounter_EndTimeStamp {
					constants.Mut.Lock()
					for _, v := range FixedWindowCounterList {
						v.Reset()
					}
					constants.Mut.Unlock()
					constants.FixedWindowCounter_EndTimeStamp = constants.FixedWindowCounter_EndTimeStamp.Add(time.Second * constants.FixedWindowCounter_WindowSize)
				}
			}
		}
	}()
}
