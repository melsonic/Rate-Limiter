package middleware

import (
	"fmt"
	"net/http"

	"github.com/melsonic/rate-limiter/algo"
	"github.com/melsonic/rate-limiter/constants"
)

var RequestTimeStamps map[string]*algo.SlidingWindowLog = make(map[string]*algo.SlidingWindowLog)

func SlidingWindowLogMiddlewareRL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClientIP := r.RemoteAddr
    constants.Mut.RLock()
		entry, entryPresent := RequestTimeStamps[ClientIP]
    constants.Mut.RUnlock()
		if !entryPresent {
			entry = &algo.SlidingWindowLog{}
      constants.Mut.Lock()
			RequestTimeStamps[ClientIP] = entry
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
