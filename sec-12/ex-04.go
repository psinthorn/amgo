package main

import "fmt"

func main() {

	//anonymouse struct
	anm := struct {
		fname string
		lname string
		//use map[key type]value type as field
		phone map[string]int

		//use slide as filed
		favFoods []string
	}{
		fname: "Sinthorn",
		lname: "Pr.",
		phone: map[string]int{
			"Nut":    555,
			"Phansa": 333,
		},
		favFoods: []string{
			"Padthai",
			"Noodle",
			"Soup",
			"Papaya Salad",
		},
	}

	fmt.Println(anm)

	fmt.Println(anm.phone)
	fmt.Printf("Type %T\n", anm.phone)
	fmt.Println(anm.favFoods)
	fmt.Printf("Type %T\n", anm.favFoods)

	//loop map with for->range key: value
	for key, value := range anm.phone {
		fmt.Println(key, value)
	}

	//loop slide composit literal by for->range
	for index, value := range anm.favFoods {
		fmt.Println(index, value)
	}

}
