package main

import "fmt"

func main() {

	// Create slide with composit literal, assing value and range and print out
	sl := [6]int{1, 3, 5, 7, 9, 11}
	fmt.Println(sl)

	//Or can create new slide without specify to number and value
	// on print out will show you an empty array
	slwo := []int{}
	fmt.Println(slwo)

	//print out length of slice
	fmt.Println(len(sl))

	//access slide by index
	fmt.Println(sl[0])

	// lists all the slide value by for loop worj with len()
	fmt.Println("Access slide value by for loop")
	for i := 0; i <= len(sl)-1; i++ {
		fmt.Println(i, ":", sl[i])
	}

	// List slide value by range over the slide using for loop
	fmt.Println("Access slide value by range")
	for i, v := range sl {
		fmt.Println(i, ":", v)
		// fmt.Printf("%T\n", sl)
	}

	// v := range sl {
	// 	fmt.Println(v)
	// }

}
