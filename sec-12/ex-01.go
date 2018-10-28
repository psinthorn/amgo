package main

import "fmt"

type person struct {
	fname   string
	lname   string
	age     int
	ice_fav []string
}

type foods struct {
	breakfast []string
	lunch     []string
	dinner    []string
	drink     []string
}

type personal_info struct {
	id int
	person
	foods
}

func main() {

	p1 := person{
		fname:   "Sin",
		lname:   "P.",
		age:     38,
		ice_fav: []string{"Coconut Milk", "Chocolate"},
	}

	fmt.Println(p1)
	for i, v := range p1.ice_fav {
		fmt.Println(i, ":", v)
	}

	p_info := personal_info{
		id: 001,
		person: person{
			fname:   "Sinthorn",
			lname:   "Pr.",
			age:     38,
			ice_fav: []string{"Coconut Milk", "Chocolate"},
		},
		foods: foods{
			breakfast: []string{"Coffee", "Hotdog"},
			lunch:     []string{"coffee", "Noodle"},
			dinner:    []string{"Milk", "Salad"},
		},
	}

	fmt.Println(p_info)

	m := map[string]personal_info{
		p_info.fname: p_info, //key:value
	}

	fmt.Println("\n")
	fmt.Println("Print map key and value")
	fmt.Println("\n")
	fmt.Println(m)
	for key, value := range m {
		fmt.Println(key, ":", value)
		fmt.Println("\n")

		//access individual value
		fmt.Println(value.fname, value.lname)
		fmt.Println("\n")

		//loop data by range
		for index, value := range value.breakfast {
			fmt.Println(index, ":", value)
		}

	}

	fmt.Println("\n")

}
