package main

import "fmt"

func main() {

	//assign anonymous func to variable
	f := func(x int, y int) int {
		fmt.Println("Argument is ", x, "and", y)
		return x * y
	}

	g := f(3, 5)
	fmt.Println("x*y=", g)

}
