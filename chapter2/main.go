// command and control

// if statement
package main

import (
	"errors"
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
// func main() {
// 	input := 10
// 	// input := -11
// 	// input := 15

// 	if input < 0 {
// 		fmt.Println("input cannot be negative")
// 	} else if input%2 == 1 {
// 		fmt.Println(input, "is odd")
// 	} else {
// 		fmt.Println(input, "is even")
// 	}
// }

// initial if statement
func validate(input int) error {
	if input < 0 {
		return errors.New("input cannot be a negative number")
	} else if input > 100 {
		return errors.New("input cannot be over 100")
	} else if input%7 == 0 {
		return errors.New("input cannot be divisible by 7")
	} else {
		return nil
	}
}

func main() {
	input := 21

	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is odd")
	} else {
		fmt.Println(input, "is even")
	}
}
