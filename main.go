package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	// Map syntax one
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
	}

	// Map syntax two
	var newColors map[string]string

	// Map syntax three
	newNewColors := make(map[string]string)
	newNewColors["black"] = "#ffff"
	delete(newNewColors, "black")

	fmt.Println(colors)
	fmt.Println(newColors)
	fmt.Println(newNewColors)
}

// Structs

// type contactInfo struct {
// 	email   string
// 	zipCode int
// }

// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }

//inside main {
// contactInfo := contactInfo{email: "yi@mail.com", zipCode: 57170}

// 	// Syntax One
// 	yiro := person{"Yiro", "Yi", contactInfo}

// 	// Syntax Two
// 	yujin := person{firstName: "Yujin", lastName: "Cho", contactInfo: contactInfo}

// 	// Syntax Three
// 	var kripto person
// 	kripto.firstName = "Kripto"
// 	kripto.lastName = "Yi"
// 	kripto.contactInfo = contactInfo

// 	yiro.updateName("Yiro Yuzin")
// 	yiro.print()
// 	yujin.print()
// 	kripto.print()
//}

// func (pointerToPerson *person) updateName(newFirstName string) {
// 	(*pointerToPerson).firstName = newFirstName
// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }
