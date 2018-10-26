package main

import "fmt"

const (
	a = 2015 + iota
	b = 2015 + iota
	c = 2015 + iota
	d = 2015 + iota
)

func main() {

	fmt.Println("this is use iota increament ", a, b, c, d)
}
