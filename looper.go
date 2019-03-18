package main

import "fmt"

func myloop(value int) {
	//print sum of all numbers till value.
	sum := 0
	for i := 0; i <= value; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	myloop(10)
}
