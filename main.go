package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	contactInfo := contactInfo{email: "yi@mail.com", zipCode: 57170}

	// Syntax One
	yiro := person{"Yiro", "Yi", contactInfo}

	// Syntax Two
	yujin := person{firstName: "Yujin", lastName: "Cho", contactInfo: contactInfo}

	// Syntax Three
	var kripto person
	kripto.firstName = "Kripto"
	kripto.lastName = "Yi"
	kripto.contactInfo = contactInfo

	yiroPointer := &yiro
	yiroPointer.updateName("Yiro Yuzin")
	yiro.print()
	yujin.print()
	kripto.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
