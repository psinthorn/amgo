package main

import "fmt"

func main() {

	myMap := map[string]int{
		"Sinthorn": 38,
		"Nut":      13,
		"Phansa":   3,
		"Pornphan": 28,
	}

	//Print out full key and value that store on map
	fmt.Println("\n")
	fmt.Println("Print all value from map")
	fmt.Println(myMap)

	//Print out individual value
	fmt.Println("\n")
	fmt.Println("Access and print all value from map")
	fmt.Println("Nut age is -> ", myMap["Nut"])
	fmt.Println("Phansa age is -> ", myMap["Phansa"])

	//below will return you a value and bool, if key found will return value and true, if no key found will return zero value and false
	fmt.Println("\n")
	fmt.Println("access key and valid check on map")
	v, ok := myMap["Phansa"]
	fmt.Println(v)
	fmt.Println(ok)

	//use if else to check is valid? and return the value
	if v, ok := myMap["Phansa1"]; ok {
		fmt.Println("Your age is -> ", v)
	} else {
		fmt.Println("No key found")
	}

	//Add new key:value to map
	myMap["Mike"] = 35

	//Use rang to pair key:value and print out
	fmt.Println("\n")
	fmt.Println("Range key and value from map")
	for key, value := range myMap {

		fmt.Println(key, ":", value)
	}

	//Range a slide index and value and print out
	xi := []int{1, 3, 5, 7, 9, 11}

	fmt.Println("\n")
	fmt.Println("Range index and value from slide")
	for index, value := range xi {

		fmt.Println(index, ":", value)
	}

	// Delete key on map
	delete(myMap, "Sinthorn")
	fmt.Println(myMap)

	// Check key befor delete if key not exist then added instead
	if v, ok := myMap["Nut"]; ok {
		fmt.Println("On process of delete -> ", v)
		delete(myMap, "Nut")
		fmt.Println("--------------->")
		fmt.Println(v, "is deleted")
		fmt.Println(myMap)
	}
}
