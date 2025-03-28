package main

import (
	"fmt"
	"time"
)

var (
	debug       bool
	loglevel    = "infotest"
	startuptime = time.Now()
)

func main() {
	fmt.Println(debug, loglevel, startuptime)
}
