package main

// function declarations
// camelCase means functions are not exportable
// but CamelCase means it can be exported

// checkNum function
import "fmt"

func checkNum() {
	for i := 1; i <= 30; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}

func main() {
	checkNum()
}
