package main

import (
	"fmt"
)

func main() {

	bd := 1980
	for {
		if bd > 2018 {
			break
		}
		fmt.Println(bd)
		bd++

	}
}
