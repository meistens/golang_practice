package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	quote := quote.Glass()

	fmt.Printf("%s\n", id)
	fmt.Printf("%s", quote)
}
