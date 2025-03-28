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
