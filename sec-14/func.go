package main

import "fmt"

//Football team struct requirement
type team struct {
	name   string
	slogan string
	desc   string
}

type contact struct {
	title    string
	addr1    string
	addr2    string
	addr3    string
	city     string
	country  string
	postCode int
	phone1   int
	mobile   int
	fax      int
	www      string
	email    string
}

type social struct {
	facebook  string
	instagram string
	youtube   string
	twitter   string
}

type teamInfo struct {
	team    //embeded struct
	contact //embeded struct
	social  //embeded struct
}

func main() {

	x := foo("Sinthorn")
	fmt.Println(x)

}

//func foo(parameter) (return type) { code }
func foo(str string) bool {

	if str == " " {
		return false
	} else {
		x := true
		fmt.Println("this string parameter", str)
		return x
	}

}
