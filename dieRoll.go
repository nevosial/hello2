package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func main() {
	fmt.Println("Hi")
	fmt.Println(dieRoll(6))
}
