package main

// errors

// runtime errors
import "fmt"

func main() {
	nums := []int{2, 4, 6, 8}

	total := 0

	for i := 0; i <= 30; i++ {
		total += nums[i]
	}
	fmt.Println("total: ", total)
}
