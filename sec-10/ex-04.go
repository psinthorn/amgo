package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := []int{13, 14, 15, 16, 17, 18}

	// append single value to slide
	sl = append(sl, 10)
	fmt.Println(sl)

	//append multi value to current slide
	sl = append(sl, 19, 20, 21)
	fmt.Println(sl)

	// Append elements to current slide use ... after elements (append is buitin func)
	sl = append(sl, x...)
	fmt.Println(sl)

}
