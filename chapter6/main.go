package main

// errors

// runtime errors
// import "fmt"

// func main() {
// 	nums := []int{2, 4, 6, 8}

// 	total := 0

// 	for i := 0; i <= 30; i++ { // use range instead to avoid this mistake
// 		total += nums[i]
// 	}
// 	fmt.Println("total: ", total)
// }

// semantic errors
import "fmt"

func main() {
	km := 2

	if km > 2 { // fix by adding a '='
		fmt.Println("take the car")
	} else {
		fmt.Println("walk")
	}
}
