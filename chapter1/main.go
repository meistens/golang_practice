package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func maintest() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}
