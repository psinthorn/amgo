package main

import (
	"fmt"
)

// Declare global variable here
var greet = "Hello"

// Declare "z" with int type if not assign value the default will be 0
var z int

// Declare "fName" variable with string type if not assign value the default value is "" (empty string)
var fName = "Sinthorn"
var lName = "Pradutnam"
var Fullname string = `"Sinthorn" <> 'Pradutnam'`

func main() {

	// # := use for declare and asign value to variable this variable will be visible only in func main()
	x := 38
	// = assign data to variable that has been declare already
	x = 83
	y := 8
	fmt.Println(x)
	fmt.Println(y)

	// fmt.Println() can return byt of value that we input and error
	n, err := fmt.Println(greet)

	fmt.Println("Byte count Println of ", greet, "is", n)
	fmt.Println("Println error return: ", err)
	Greeting()
	PrintFullname()
	fmt.Printf("%T\n", Fullname)
	fmt.Printf("%T\n", x)

}

func Greeting() {
	fmt.Println(greet, fName)
}

func PrintFullname() {
	fmt.Println(Fullname)
}
