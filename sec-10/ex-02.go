package main

import "fmt"

func main() {
	fmt.Println("Slide and range exercise.\n")

	// create slide of int with 10 value
	sl := []int{2, 4, 6, 8, 10, 11, 12, 13, 14, 15}

	for i, v := range sl {
		fmt.Println(i, ":", v)
	}
	fmt.Printf("Type of this slide is %T\n", sl)
}
