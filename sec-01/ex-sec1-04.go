package main

import "fmt"

// Create new type
type goint int

var x goint

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

}
