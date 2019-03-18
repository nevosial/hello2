package main

import "fmt"

func main() {
	//array
	a := [7]int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println(a)

	//slice
	b := a[1:3]
	fmt.Println(b)

	//a slice is an array literal without a length.
	var c = []bool{true, false, true}
	fmt.Println(c)

	//changing the value in the slice will change the value
	//in the array that it references.
	b[0] = 99
	fmt.Println(b)
	fmt.Println(a)

	//slice of structs. [{},{},{}]
	var sos = []struct {
		name string
		age  int
	}{
		{"nev", 20},
		{"nevi", 30},
		{"dnev", 40},
		{"vnev", 50},
	}

	fmt.Println(sos)
	fmt.Println(sos[1])
	fmt.Println(sos[1].name)
	fmt.Println(sos[1].age)

	//length of a slice : len(s) gives the size of the current slice s.
	fmt.Println(len(b))

	//capacity of the slice : cap(s) gives the size of the underlying array.
	fmt.Println(cap(b))
}
