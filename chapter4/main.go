package main

// defining an array
// import (
// 	"fmt"
// )

// func defineArr() [10]int {
// 	var arr [10]int
// 	return arr
// }
// func main() {
// 	fmt.Printf("%#v\n", defineArr())
// }

// init. an arr. using keys
// import "fmt"

// func compArr() (bool, bool, [10]int) {
// 	arr1 := [10]int{}
// 	// set key 9 to value 0
// 	arr2 := [...]int{9: 0}
// 	// set key 0 to value 1, set key 9 to value 10, set key 4 to value 5
// 	arr3 := [10]int{1, 9: 10, 4: 5}

// 	return arr1 == arr2, arr1 == arr3, arr3
// }

// func main() {
// 	comp1, comp2, arr3 := compArr()
// 	fmt.Println("[10]int == [...]{9:0}	:", comp1)
// 	fmt.Println("[10]int==[10]int{1, 9:10, 4:5}	:", comp2)
// 	fmt.Println(arr3)
// }

// reading a single item from an array
// import "fmt"

// func msg() string {
// 	arr := [...]string{
// 		"ready",
// 		"set",
// 		"go",
// 		"testing",
// 		"it works",
// 	}
// 	return fmt.Sprintln(arr[0], arr[2], arr[1], arr[4])
// }

// func main() {
// 	fmt.Println(msg())
// }

// writing to ann arr
// import "fmt"

// func msg() string {
// 	arr := [4]string{"is", "this", "working", "???"}

// 	arr[1] = "that"
// 	arr[2] = "it..."

// 	return fmt.Sprintln(arr[0], arr[1], arr[2], arr[3])
// }

// func main() {
// 	fmt.Println(msg())
// }

// looping over an arr
// import "fmt"

// func msg() string {
// 	m := ""
// 	arr := [4]int{1, 2, 3, 4}

// 	for i := range arr {
// 		// multiply
// 		arr[i] = arr[i] * arr[i]
// 		// add result to the empty variable m
// 		m += fmt.Sprintf("%v: %v\n", i, arr[i])
// 	}
// 	return m
// }

// func main() {
// 	fmt.Printf(msg())
// }

// modifying the contents of an array in a loop
// import "fmt"

// func fillArr(arr [10]int) [10]int {
// 	//using range where necessary
// 	//not bothering with saving computing cost or whatever right now
// 	//plus, shouldn't computing hardware be too powerful, unless
// 	// there are many still using toasters as hardware...
// 	for i := range arr {
// 		//increment the value of the provided key(i) by 1 till array is filled to 10
// 		arr[i] = i + 1
// 	}
// 	return arr
// }

// // multiply number from an array by itself and sets the result back to the array
// func multArr(arr [10]int) [10]int {
// 	for i := range arr {
// 		arr[i] = arr[i] * arr[i]
// 	}
// 	return arr
// }

// func main() {
// 	var arr [10]int
// 	arr = fillArr(arr)
// 	arr = multArr(arr)
// 	fmt.Println(arr)
// }

// nil value
// import "fmt"

// func main() {
// 	var msg []string
// 	if msg == nil {
// 		fmt.Println("error, unexpected nil value")
// 		return
// 	}
// 	fmt.Println(msg)
// }

// filling an array
// import "fmt"

// func fillArr(arr [10]int) [10]int {
// 	// using range to create and fill an array from 1 to 10
// 	for i := range arr {
// 		arr[i] = i + 1
// 	}
// 	return arr
// }

// func main() {
// 	// create a variable of array to hold the created array above
// 	var arr [10]int
// 	arr = fillArr(arr)
// 	fmt.Println(arr)
// }

// slices
// import (
// 	"fmt"
// 	"os"
// )

// // return value is a str. slice
// func passedArgs(minArgs int) []string {
// 	if len(os.Args) < minArgs {
// 		fmt.Printf("at least %v arguments are needed\n", minArgs)
// 		os.Exit(1)
// 	}
// 	var args []string
// 	for i := 1; i < len(os.Args); i++ {
// 		args = append(args, os.Args[i])
// 	}
// 	return args
// }

// func finlongest(args []string) string {
// 	var longest string

// 	for i := range args {
// 		if len(args[i]) > len(longest) {
// 			longest = args[i]
// 		}
// 	}
// 	return longest
// }

// func main() {
// 	if longest := finlongest(passedArgs(3)); len(longest) > 0 {
// 		fmt.Println("longest passed:", longest)
// 	} else {
// 		fmt.Println("error")
// 		os.Exit(0)
// 	}
// }

// creating slices from a slice
// import "fmt"

// func msg() string {
// 	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	m := fmt.Sprintln(s[0], s[0:1], s[:1])

// 	m += fmt.Sprintln(s[len(s)-1], s[len(s)-1:len(s)])

// 	return m
// }

// func main() {
// 	fmt.Print(msg())
// }
