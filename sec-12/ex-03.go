package main

import "fmt"

// Create a struct here
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truckNew := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	fmt.Println("Truck Info")
	fmt.Println(truckNew)
	fmt.Println(truckNew.doors)
	fmt.Println(truckNew.color)

	sedanNew := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println("Sedan Info")
	fmt.Println(sedanNew)
	fmt.Println("Doors", sedanNew.doors)
	fmt.Println("Color", sedanNew.color)

}
