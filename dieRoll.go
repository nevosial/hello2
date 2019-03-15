package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func onedieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func twodieRoll(size1, size2 int) (int, int) {
	return onedieRoll(size1), onedieRoll(size2)
}

func returnsNamed(ip1 string, ip2 int) (theresult string, err error) {
	theresult = "modified " + ip1 + " , " + strconv.Itoa(ip2)
	return theresult, err
}

func main() {
	fmt.Println("Hi")
	fmt.Println(onedieRoll(20))
	fmt.Println(twodieRoll(25, 35))

	named, err := returnsNamed("globule", 42)
	fmt.Printf("Named params returned: '%s', %v\n", named, err)

}
