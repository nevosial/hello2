package main

import (
	"fmt"
	"runtime"
	"time"
)

func printeven() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func getprice(a string) int {
	switch a {
	case "apple":
		return 10
	case "bun":
		return 50
	default:
		return 0
	}
}

func findOS() string {
	os := runtime.GOOS
	switch os {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	default:
		return os
	}
}

//switch without condition
func swoc() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning !")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon !")
	default:
		fmt.Println("Good Evening !")
	}
}

func main() {
	printeven()
	println(getprice("apple"))
	println(getprice("bun"))
	println(getprice("kela"))
	print("Go runs on: ")
	println(findOS())
	swoc()
}
