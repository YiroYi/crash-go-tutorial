package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "Is working correct")
	c <- link
}

// Cards
// cards := newDeck()
// cards.shuffle()
// cards.print()

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

// Maps

// Map syntax two
// var newColors map[string]string

// Map syntax three
// newNewColors := make(map[string]string)
// newNewColors["black"] = "#ffff"
// delete(newNewColors, "black")

// fmt.Println(colors)
// fmt.Println(newColors)
// fmt.Println(newNewColors)
// }

// Map syntax one
// colors := map[string]string{
// 	"red":   "#ff000",
// 	"green": "#4bf745",
// 	"white": "#4ffff",
// }

// printMap(colors)
// fmt.Println(colors)
// }

// func printMap(c map[string]string) {
// //for key, value := range c {
// for color, hex := range c {
// 	fmt.Println(" Hex code for", color, "is", hex)
// }

// Interface
// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct{}
// type spanishBot struct{}

// func main() {
// 	eb := englishBot{}
// 	sb := spanishBot{}

// 	printGreeting(eb)
// 	printGreeting(sb)
// }

// func printGreeting(b bot) {
// 	fmt.Println(b.getGreeting())
// }

// func (englishBot) getGreeting() string {
// 	return "Hi There"
// }

// func (spanishBot) getGreeting() string {
// 	return "Hola"
// }

//func main() {
// 	resp, err := http.Get("http://google.com")

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// 	io.Copy(os.Stdout, resp.Body)
// }
