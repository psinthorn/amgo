package main

import "fmt"

type gostring int

var x gostring
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 88
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// conversion
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
