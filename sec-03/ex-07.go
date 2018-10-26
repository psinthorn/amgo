package main

import "fmt"

func main() {
	x := "Jame Bond"
	// y := "Sinthorn Pr."

	fmt.Println("All th below will work by else if and else\n")

	//If on of condition is true program will do the statement and exit imediatly
	if x == "Jame Bond" {
		fmt.Println("Bond, Jame Bond")
	} else if x == "Sinthorn P." {
		fmt.Println("Sinthorn? Yes It's me")
	} else {
		fmt.Println("Who are you?")
	}
}
