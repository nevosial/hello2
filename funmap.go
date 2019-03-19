//Maps in Go
package main

import "fmt"

type Student struct {
	wt  int
	age int
}

var m = map[string]Student{
	"nev": {100, 20},
	"dan": {80, 25},
}

func main() {
	fmt.Println(m)
	fmt.Println(m["dan"])
	fmt.Println(m["nev"])

	user := map[string]string{
		"name": "Ken",
		"age":  "20cms",
	}

	height, ok := user["height"]
	if ok == true {
		fmt.Println("Height: ", height)
	} else {
		fmt.Println("Height not found.")
	}
}
