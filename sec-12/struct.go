package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
	code  int
}

type secretAgent struct {
	person
	code int
	ltk  bool
}

func main() {

	p1 := person{
		fname: "Mike",
		lname: "P.",
		age:   38,
		code:  1008,
	}

	sa1 := secretAgent{
		person: person{
			fname: "Mike",
			lname: "P.",
			age:   38,
			code:  1003, //collistion key
		},
		code: 1008, //"code":key will be collision with person embeded inner struct above but you still can work together see below.
		ltk:  true,
	}

	//Print out full deatial of p1
	fmt.Println(p1)

	//Print out full detail of sa1
	fmt.Println(sa1)

	//Access and print out individual key:value of sa1.
	//if embeded inner struct key is not collision with main struct key the struct will promote inner type key to outer struct type key.
	//and access embeded inner type that collision with outer type as below
	// ***Summary When you use another type struct in another struct we call "Embeded type" embeded type will promate to outer type on current struct

	fmt.Println("SA1 ", sa1.fname, sa1.lname, sa1.age, sa1.person.code, sa1.code, sa1.ltk)

}
