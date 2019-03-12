package main

import "fmt"

func printeven(){
	for i:=0; i<=10; i++{
		if i%2==0 {
			fmt.Println(i)
		}
	}
}

func getprice(a string) int {
	switch a{
		case "apple":
			return 10
		case "bun":
			return 50
		default: 
			return 0
	}
}


func main(){
	printeven()
	println(getprice("apple"))
	println(getprice("bun"))
	println(getprice("kela"))

}

