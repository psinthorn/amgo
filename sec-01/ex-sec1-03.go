package main

import (
	"fmt"
)

var x int = 42
var y string = "jame Bond"
var z bool = true

func main() {

	// Use Printf to print out type of each value in single line
	fmt.Printf("%T\n%T\n%T\n", x, y, z)

	// Use short decalre variavle "s" and assign value of x y z to s variable then print that value out
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)

}
