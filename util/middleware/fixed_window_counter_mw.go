package middleware

import (
	"fmt"
	"net/http"

	"github.com/melsonic/rate-limiter/algo"
	"github.com/melsonic/rate-limiter/constants"
)

var FixedWindowCounterList map[string]*algo.FixedWindowEntry = make(map[string]*algo.FixedWindowEntry)

func FixedWindowCounterMiddlewareRL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClientIP := r.RemoteAddr
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
