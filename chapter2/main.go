// command and control

// if statement
package main

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
// func validate(input int) error {
// 	if input < 0 {
// 		return errors.New("input cannot be a negative number")
// 	} else if input > 100 {
// 		return errors.New("input cannot be over 100")
// 	} else if input%7 == 0 {
// 		return errors.New("input cannot be divisible by 7")
// 	} else {
// 		return nil
// 	}
// }

// func main() {
// 	input := 21

// 	if err := validate(input); err != nil {
// 		fmt.Println(err)
// 	} else if input%2 == 0 {
// 		fmt.Println(input, "is odd")
// 	} else {
// 		fmt.Println(input, "is even")
// 	}
// }

// expression switch statements

// func main() {
// 	dayBorn := time.Monday

// 	switch dayBorn {
// 	case time.Monday:
// 		fmt.Println("Monday's child is fair of face")
// 	case time.Tuesday:
// 		fmt.Println("Tuesday's child is full of grace")
// 	// this goes on and on...
// 	default:
// 		fmt.Println("Error, day born is not valid")
// 	}
// }

// switch and multiple case values
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	dayBorn := time.Tuesday

// 	switch dayBorn {
// 	case time.Wednesday, time.Thursday, time.Friday:
// 		fmt.Println("born on a weekday")
// 		// skipped in case of error, dunno how to handle that...yet
// 	default:
// 		fmt.Println("born on some weekday")
// 	}
// }

// expressionless switch statements
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	switch dayBorn := time.Sunday; {
// 	case dayBorn == time.Sunday || dayBorn == time.Saturday:
// 		fmt.Println("born on a weekend")
// 	default:
// 		fmt.Println("born on some other day")
// 	}
// }

// loops
// import (
// 	"fmt"
// )

// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// looping over array nd slices
// import (
// 	"fmt"
// )

// func main() {
// 	names := []string{"Stewie", "Megatron", "Peter", "Giggity"}

// 	for i := 0; i < len(names); i++ {
// 		fmt.Println(names[i])
// 	}
// }

// range loops over a map
import (
	"fmt"
)

func main() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}
