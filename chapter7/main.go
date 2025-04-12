package main

//interfaces
import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
}

func main() {
	c := cat{}
	fmt.Println(c.Speak())
	c.Greeting()
}

func (c cat) Speak() string {
	return "meow"
}

func (c cat) Greeting() {
	fmt.Println("meow?????")
}
