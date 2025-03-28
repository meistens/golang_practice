package main

import (
	"fmt"
	"time"
)

var foo string = "test"

func maintest2() {
	var bar string = "another test"

	fmt.Println(foo, bar)
}

var (
	debug       bool      = false
	loglevel    string    = "info"
	startupTime time.Time = time.Now()
)

func main() {
	fmt.Println(debug, loglevel, startupTime)
}
