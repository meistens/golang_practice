// core types
package main

// program to measure password complexity
// import (
// 	"fmt"
// 	"unicode"
// )

// func pwdChecker(pw string) bool {
// 	// checkout rune, add to stuff_to_know_about
// 	pwdR := []rune(pw)

// 	if len(pwdR) < 8 {
// 		return false
// 	}

// 	hasUpper := false
// 	hasLower := false
// 	hasNumber := false
// 	hasSymbol := false

// 	// '_' is a catch-all, in this case for the key
// 	for _, value := range pwdR {
// 		if unicode.IsUpper(value) {
// 			hasUpper = true
// 		}
// 		if unicode.IsLower(value) {
// 			hasLower = true
// 		}
// 		if unicode.IsNumber(value) {
// 			hasNumber = true
// 		}
// 		if unicode.IsPunct(value) || unicode.IsSymbol(value) {
// 			hasSymbol = true
// 		}
// 	}
// 	return hasUpper && hasLower && hasNumber && hasSymbol
// }

// func main() {
// 	if pwdChecker("yyy9wds.A") {
// 		fmt.Println("pwd good")
// 	} else {
// 		fmt.Println("pwd bad")
// 	}
// }

// overflow and wraparound
import "fmt"

// func main() {
// 	var a int8 = 127 // 128
// 	fmt.Println(a)
// }

func main() {
	var a int8 = 125
	var b uint8 = 254

	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println("int8: ", a, "uint8: ", b)
	}
}
