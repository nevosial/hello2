package main

import (
	"fmt"
	"math/rand"
	"time"
)

func onedieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func twodieRoll(size1, size2 int) (int, int) {
	return onedieRoll(size1), onedieRoll(size2)
}

func main() {
	fmt.Println("Hi")
	fmt.Println(onedieRoll(20))
	fmt.Println(twodieRoll(25, 35))
}
