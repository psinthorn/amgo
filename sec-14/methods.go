package main

import "fmt"

// Struct type
type person struct {
	fname  string
	lname  string
	gender string
	age    int
}

// struct and embeded struct
type playerInfo struct {
	person
	no       int
	position []string          //slide of string
	foot     map[string]string //map key:string and value:string
}

func main() {

	player1 := playerInfo{
		person: person{
			fname:  "Sinthorn",
			lname:  "Pradutnam",
			gender: "Male",
			age:    38,
		},
		no:       18,
		position: []string{"MD", "CB"},
		foot:     map[string]string{"primary": "right"},
	}

	player2 := playerInfo{
		person: person{
			fname:  "Nut",
			lname:  "Pradutnam",
			gender: "Male",
			age:    13,
		},
		no:       1,
		position: []string{"GK"},
		foot:     map[string]string{"primary": "right"},
	}

	fmt.Println(player1)
	player1.getInfo()
	player2.getInfo()

}

// create func and method and then attachment this method to each player
// func (r rceiver) idntifier(parameter(s)) (return(s)) {}
func (player playerInfo) getInfo() {
	fmt.Println("Player: ", player.fname, player.lname, player.no, player.position[0], player.foot["primary"])
}
