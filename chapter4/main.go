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
import "fmt"

func msg() string {
	arr := [...]string{
		"ready",
		"set",
		"go",
		"testing",
		"it works",
	}
	return fmt.Sprintln(arr[0], arr[2], arr[1], arr[4])
}

func main() {
	fmt.Println(msg())
}
