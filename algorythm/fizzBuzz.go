package main

import "fmt"

func main() {
	//FizzBuzz
	fizzBuzz(50)

}

//FizzBuzz Func
func fizzBuzz(num int) {
	fmt.Println("fizzBuzz Algorythm")
	for i := 1; i <= num; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
