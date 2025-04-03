package main

// defining an array
import (
	"fmt"
)

func defineArr() [10]int {
	var arr [10]int
	return arr
}
func main() {
	fmt.Printf("%#v\n", defineArr())
}
