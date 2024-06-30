package main

import (
	"fmt"
	"net/http"

	"github.com/melsonic/rate-limiter/constants"
	"github.com/melsonic/rate-limiter/util"
)

func main() {
  fmt.Println(constants.ServerAddr)
	serverErr := http.ListenAndServe(constants.ServerAddr, util.Route())
	if serverErr != nil {
		fmt.Printf("Server crashed\n")
	}
}
