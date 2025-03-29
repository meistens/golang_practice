// command and control

// if statement
package main

import (
	"fmt"
)

// func main() {
// 	input := 8

// 	if input%2 == 0 {
// 		fmt.Println(input, "is even")
// 	}

// 	if input%2 == 1 {
// 		fmt.Println(input, "is odd")
// 	}
// }

// if else
// func main() {
// 	input := 8

// 	if input%2 == 0 {
// 		fmt.Println(input, "is even")
// 	} else {
// 		fmt.Println(input, "is odd")
// 	}
// }

// else if
func main() {
	input := 10
	// input := -11
	// input := 15

	if input < 0 {
		fmt.Println("input cannot be negative")
	} else if input%2 == 1 {
		fmt.Println(input, "is odd")
	} else {
		fmt.Println(input, "is even")
	}
}
