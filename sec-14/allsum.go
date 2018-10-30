package main

import "fmt"

func main() {

	//slide of int to callect all the number
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// pass slide of int as arguments to sum func
	grandTotal := sum(ii...)

	//Print out the result from call back of sum func
	fmt.Println(grandTotal)
}

// receive argumensts to sum parametsrs
func sum(numLists ...int) int {

	// assign total variable
	total := 0

	// use for loop to range out all the value of slide
	for _, v := range numLists {

		// use += to sum all the value and assign to total variable
		total += v
	}

	// return tatal to function
	return total
}
