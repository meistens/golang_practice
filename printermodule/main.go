package main

import (
	"fmt"

	"github.com/sicoyle/printer"
)

func main() {
	msg := printer.PrintNewUID()
	fmt.Println(msg)
}
