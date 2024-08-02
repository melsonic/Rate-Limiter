package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/melsonic/rate-limiter/algo"
	"github.com/melsonic/rate-limiter/constants"
)

var SlidingWindowCounterEntryList map[string]*algo.SlidingWindowCounterEntry = make(map[string]*algo.SlidingWindowCounterEntry)

func SlidingWindowCounterMiddlewareRL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClientIP := strings.Split(r.Header.Get("X-FORWARDED-FOR"), ", ")[0]
		if ClientIP == "" {
			ClientIP = r.RemoteAddr
		}
		constants.Mut.RLock()
		entry, entryPresent := SlidingWindowCounterEntryList[ClientIP]
		constants.Mut.RUnlock()
		if !entryPresent {
			var windowStart time.Time = constants.SlidingWindowCounter_WindowStartTime
			entry = &algo.SlidingWindowCounterEntry{PrevWindowCount: 0, CurrentWindowCount: 0, WindowStartTime: windowStart, WindowEndTime: windowStart.Add(time.Second * constants.SlidingWindowCounter_WindowSize)}
			constants.Mut.Lock()
			SlidingWindowCounterEntryList[ClientIP] = entry
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

func SlidingWindowCounterHelper() {
	go func() {
		var ticker *time.Ticker = time.NewTicker(time.Second * constants.SlidingWindowCounter_WindowSize)
		for {
			constants.SlidingWindowCounter_WindowStartTime = time.Now()
			constants.Mut.Lock()
			for _, v := range SlidingWindowCounterEntryList {
				v.WindowReset()
			}
			constants.Mut.Unlock()
			<-ticker.C
		}
	}()
}
