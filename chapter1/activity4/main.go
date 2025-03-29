package main

// import "math/rand"

// func main() {
// 	var seed int64 = 1234567
// 	rand.NewSource(seed)
// }

// short variable declaration (inside the function body)
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	debug := false
// 	loglevel := "innnffffooo"
// 	startuptime := time.Now()

// 	fmt.Println(debug, loglevel, startuptime)
// }

// ***declaring multiple variables with a short variable declaration
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	debug, loglevel, startuptime := false, "info", time.Now()
// 	fmt.Println(debug, loglevel, startuptime)
// }

// declaring multiple vars from a func
// import (
// 	"fmt"
// 	"time"
// )

// func getConfig() (bool, string, time.Time) {
// 	return false, "ugh...", time.Now()
// }

// func main() {
// 	debug, loglevel, starttime := getConfig()
// 	fmt.Println(debug, starttime, loglevel)
// }

// using var to declare multiple variables in one line
// import (
// 	"fmt"
// 	"time"
// )

// func getConfig() (bool, string, time.Time) {
// 	return false, "info", time.Now()
// }

// func main() {
// 	// type only
// 	var start, middle, end float32
// 	fmt.Println(start, middle, end)

// 	// initial value mixed types
// 	var name, left, right, top, bottom = "one", 1, 1.3, 2.0, 3
// 	fmt.Println(name, left, right, top, bottom)

// 	// works with functions also, but prefer :=
// 	var debug, loglevel, startuptime = getConfig()
// 	fmt.Println(debug, loglevel, startuptime)
// }

// changing the value of an already declared variable
// import "fmt"

// func main() {
// 	// declared here
// 	offset := 65
// 	fmt.Println(offset)
// 	// changed here
// 	offset = 10
// 	fmt.Println(offset)
// }

// changing multiple values at once
// import "fmt"

// func main() {
// 	query, limit, offset := "bat", 10, 0
// 	query, limit, offset = "ball", offset, 20

// 	fmt.Println(query, limit, offset)
// }

// OPERATORS
// import "fmt"

// func main() {
// 	total := float64(2 * 13)
// 	fmt.Println("sub:", total)

// 	total = total + (4 * 2.25)
// 	fmt.Println("sub:", total)

// 	total = total - 5
// 	fmt.Println("sub:", total)

// 	tip := total * 0.1
// 	fmt.Println("tip:", tip)

// 	total = total + tip
// 	fmt.Println("total:", total)

// 	split := total / 2
// 	fmt.Println("split:", split)

// 	visitcount := 24
// 	visitcount = visitcount + 1

// 	remainder := visitcount % 5

// 	if remainder == 0 {
// 		fmt.Println("you have earned a reward with this visit and purchase")
// 	}
// }

// implement shorthand operators
// import "fmt"

// func main() {
// 	count := 5
// 	count += 5
// 	fmt.Println(count)

// 	count++
// 	fmt.Println(count)

// 	count--
// 	fmt.Println(count)

// 	count -= 5
// 	fmt.Println(count)

// 	name := "meh"
// 	name += " oof"
// 	fmt.Println("hello,", name)
// }

// comparing values
// import "fmt"

// func main() {
// 	visits := 15
// 	fmt.Println("first visit	:", visits != 1)

// 	fmt.Println("silver member	:", visits >= 10 && visits <= 21)
// 	fmt.Println("gold member	:", visits > 10 && visits <= 30)
// 	fmt.Println("platinum member	:", visits > 30)
// }

// zero values
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var count int
// 	fmt.Printf("count	: %#v \n", count)

// 	var discount float64
// 	fmt.Printf("count	: %#v \n", discount)

// 	var debug bool
// 	fmt.Printf("debug	: %#v \n", debug)

// 	var message string
// 	fmt.Printf("message	: %#v \n", message)

// 	var emails []string
// 	fmt.Printf("emails	: %#v \n", emails)

// 	var starttime time.Time
// 	fmt.Printf("start	: %#v \n", starttime)
// }

// getting a pointer
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// declare pointer using a var statement
// 	var count1 *int

// 	// create variable using new
// 	count2 := new(int)

// 	// temp. variable to hold a number bcos cannot take the address of a literal number
// 	counttemp := 5

// 	// using & to create a pointer from the existing variable (referening?)
// 	count3 := &counttemp

// 	// create pointer from a type without a temp. variable
// 	t := &time.Time{}

// 	fmt.Printf("count1:	%#v\n", count1)
// 	fmt.Printf("count2:	%#v\n", count2)
// 	fmt.Printf("count3:	%#v\n", count3)
// 	fmt.Printf("time:	%#v\n", t)
// }

// getting values from a pointer
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var count1 *int
// 	count2 := new(int)
// 	counttemp := 5
// 	count3 := &counttemp
// 	t := &time.Time{}

// 	// for count1-3, add a nil to check and add * in front of the variable name
// 	if count1 != nil {
// 		fmt.Printf("count1: %#v\n", *count1)
// 	}
// 	if count2 != nil {
// 		fmt.Printf("count2:	%#v\n", *count2)
// 	}
// 	if count3 != nil {
// 		fmt.Printf("count3:	%#v\n", *count3)
// 	}

// 	// check, but deference the variable using *
// 	if t != nil {
// 		fmt.Printf("time:	%#v\n", *t)
// 	fmt.Printf("time	: %#v\n", t.String())
// 	}
// }

// function design with pointers
import "fmt"

func add5Val(count int) {
	count += 5
	fmt.Println("add5val	:", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point	:", *count)
}

func main() {
	var count int
	add5Val(count)
	fmt.Println("add5val post:", count)

	add5Point(&count)
	fmt.Println("add5Point post:", count)
}
