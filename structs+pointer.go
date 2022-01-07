#################### MAIN.GO ################




package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// method 1
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 4000,
		},
	}

	alexPointer := &alex // & -> this will give the memory of the value of alex
	alexPointer.updateName("alexy")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
