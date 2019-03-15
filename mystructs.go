package main

import "fmt"

type power struct {
	attack  int
	defense int
}

type location struct {
	x float32
	y float32
	z float32
}

type nonplayerchar struct {
	name  string
	speed int
	power power
	loc   location
}

func main() {
	fmt.Println(".....Structs.....")

	demon := nonplayerchar{
		name:  "alfio",
		speed: 21,
		power: power{attack: 40, defense: 20},
		loc:   location{x: 0.35, y: 11.11, z: 4.5},
	}

	fmt.Println(demon)
}
