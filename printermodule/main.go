package main

import (
	"fmt"

	"github.com/sicoyle/printer"
)

// update: reading the docs this time proved to be better this time...
// not gonna change this, on another project yes, an example? nah
func main() {
	msg := printer.PrintNewUID()
	fmt.Println(msg)
}
