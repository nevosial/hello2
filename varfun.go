package main
//const n_global_variable = "im package level available"

import "fmt"

//i will add two ints and return an int
func add(a, b int) int {
	return a+b
}

//i will swap two values
func swap(c, d int) (int, int){
	//im return two values.
	return d, c
}

func main(){
	a:= "im local"
	fmt.Println(a)
	k:= add(5,1)
	fmt.Println(k)
	z,x := swap(10,20)
	fmt.Println(z)
	fmt.Println(x)
	//im a local function assigined to a local var.
	n := func(){
		fmt.Println("im assigned to variable but im a function.")
	}
	n()
	}


