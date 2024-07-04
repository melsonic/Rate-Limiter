package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/melsonic/rate-limiter/constants"
	"github.com/melsonic/rate-limiter/util"
)

func main() {
	for _, str := range constants.AlgorithmOptionsArray {
		fmt.Println(str)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Select : ")
	option, err := reader.ReadString('\n')
	if err != nil || option == "" {
		if err != nil {
			log.Fatal(err)
		} else {
			log.Fatal("Please select one option\n")
		}
	}
	option = strings.TrimSpace(option)
	serverErr := http.ListenAndServe(constants.ServerAddr, util.Route(option))
	if serverErr != nil {
		fmt.Printf("Server crashed\n")
	}
}
