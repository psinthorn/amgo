package main

import "fmt"

//Football team struct requirement
type team struct {
	name   string
	slogan string
	desc   string
	year   int
}

type contact struct {
	title    string
	addr1    string
	addr2    string
	addr3    string
	city     string
	country  string
	postCode int
	//contacNo map[string]int{ "phone": 555-555, "mobile": 888-888, "fax": 000-000 }
	phone1 int
	mobile int
	fax    int
	www    string
	email  string
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

	ft := team{
		name:   "Sin FC.",
		slogan: "Imagine and Fun ทำในสิ่งที่ตัวเองเก่งที่สุด",
		desc:   "สมัครเล่น",
		year:   1980,
	}

	fmt.Println(ft)

	order := []int{1, 3, 5, 7, 9, 11}

	x := foo("Sinthorn")
	fmt.Println(x)

	sum(order)

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

func sum(sumTotal []int) int {
	order := sumTotal

	fmt.Println(order)
	return 1
}
