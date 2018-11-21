package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Sum all num with sum()")
	grandTotal := sum(ii...)
	fmt.Println("Grand Total", grandTotal)

	fmt.Println("Sum even num with callback sum() ")
	evenGrandtotal := evenSum(sum, ii...)
	fmt.Println("Even total is : ", evenGrandtotal)

	fmt.Println("Sum odd num with callback sum() ")
	oddTotal := oddSum(sum, ii...)
	fmt.Println("Odd total is : ", oddTotal)

}

//sum function
func sum(numLists ...int) int {
	total := 0
	for _, v := range numLists {
		total += v
	}
	return total
}

// Even Tatal callback func
func evenSum(sumEven func(evenList ...int) int, numLists ...int) int {

	// this callback function will back total of even number
	var evenSlide []int
	for _, v := range numLists {
		if v%2 == 0 {
			evenSlide = append(evenSlide, v)
		}
	}
	return sumEven(evenSlide...)
}

// Odd callback func
func oddSum(sumOdd func(oddList ...int) int, numList ...int) int {
	var oddSlide []int

	// find odd number from numLists argument
	for _, v := range numList {
		if v%2 != 0 {
			oddSlide = append(oddSlide, v)
		}
	}

	return sumOdd(oddSlide...)

}
