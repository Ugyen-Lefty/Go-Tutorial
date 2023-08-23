package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email string
	zip   int
}

func main() {

	// Weird but still correct
	// name := person{"Ugyen", "Lefty"}

	// Other way
	// var name person
	// name.firstName = "Ugyen"
	// name.lastName = "Lefty"

	//Best Way IMO
	name := person{
		firstName: "Ugyen",
		lastName:  "Lefty",
		contact: contact{
			email: "ugyen@gmail.com",
			zip:   11001,
		},
	}
	// namePointer := &name
	name.updateName("Lefty")
	name.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
