package main

import "fmt"

func main() {

	//Anonymous func
	//syntex pattern func(){}()
	func() {
		fmt.Println("This is anonymous func without argument passing in")
	}()

	//Anonemous func wit passing in argument
	func(x int) {
		fmt.Println("This is func passing in argument int", x)
	}(38)

}
