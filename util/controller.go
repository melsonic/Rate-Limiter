package util

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Rate-Limiter\n")
}

func UnlimitedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yo, requests are freely unlimited\n")
}

func LimitedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requests are limited, make each request count\n")
}
