package main

import "fmt"

func main() {

	sl := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println(sl[:5])  //result should be [2,4,6,8,10]
	fmt.Println(sl[5:])  //result should be [12,14,16,18,20]
	fmt.Println(sl[1:6]) //result should be [4,6,8,10,12]
	fmt.Println(sl[2:7]) //result should be [6,8,10,12,14]
}
