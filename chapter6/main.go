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
// import "fmt"

// func main() {
// 	km := 2

// 	if km > 2 { // fix by adding a '='
// 		fmt.Println("take the car")
// 	} else {
// 		fmt.Println("walk")
// 	}
// }

//
// import (
// 	"errors"
// 	"fmt"
// )

// var (
// 	ErrHrRate    = errors.New("invalid hourly rate")
// 	ErrHrsWorked = errors.New("invalid hours worked per week")
// )

// func payDay(hrsWorked, hrRate int) (int, error) {
// 	if hrsWorked < 0 || hrsWorked > 80 {
// 		return 0, ErrHrsWorked
// 	}

// 	if hrRate < 10 || hrRate > 75 {
// 		return 0, ErrHrRate
// 	}

// 	if hrsWorked > 40 {
// 		hrsOver := hrsWorked - 40
// 		oT := hrsOver * 2
// 		regPay := hrsWorked * hrRate
// 		return regPay + oT, nil
// 	}
// 	return hrsWorked * hrRate, nil
// }

// func main() {
// 	pay, err := payDay(81, 50)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	pay, err = payDay(80, 5)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	pay, err = payDay(80, 50)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(pay)
// }

// panic
// import "fmt"

// func main() {
// 	nums := []int{1, 2, 4}

// 	for i := range 10 { // fix by setting correct range to match
// 		fmt.Println(nums[i])
// 	}
// }

// basic use of panic
// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	msg := "bye"
// 	message(msg)
// 	fmt.Println("line will not be printed")
// }

// func message(msg string) {
// 	if msg == "bye" {
// 		panic(errors.New("something went wrong"))
// 	}
// }
