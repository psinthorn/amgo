package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//need to delete 4,5,6 out and new slice will be sl[1,2,3,7,8,9]
	// solution1
	sl = append(sl[:3], sl[6:]...)
	fmt.Println(sl)

}
