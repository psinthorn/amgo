package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {

	// Print default value for variable that declare without assign value
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("\n")
	fmt.Print(`Call of the decalre variable that not assign the value is call "ZERO Value"`)
	fmt.Println("\n")

	// Print type of variable
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

}
