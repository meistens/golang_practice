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

// func main() {
// 	var a int8 = 127 // 128
// 	fmt.Println(a)
// }

// func main() {
// 	var a int8 = 125
// 	var b uint8 = 254

// 	for range 5 {
// 		a++
// 		b++
// 		fmt.Println("int8: ", a, "uint8: ", b)
// 	}
// }

// big numbers
// import (
// 	"fmt"
// 	"math"
// 	"math/big"
// )

// func main() {
// 	intA := math.MaxInt64
// 	intA = intA + 1

// 	bigA := big.NewInt(math.MaxInt64)
// 	bigA.Add(bigA, big.NewInt(1))

// 	fmt.Println("MaxIt64:	", math.MaxInt64)
// 	fmt.Println("Int:	", intA)
// 	fmt.Println("Big int:	", bigA.String())
// }

// rune(s)
import "fmt"

func main() {
	uname := "Some_NAme_a+-eeU"

	for i := range len(uname) {
		fmt.Println(uname[i], " ")
	}
}
