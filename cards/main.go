package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//cards := newDeck()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("card.txt")
	//cards := newDeckFronFile("card.txts")
	//cards.print()

	//ards := newDeck()
	//cards.shuffle()
	//cards.print()

	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "94000",
		},
	}

	jim.updateName("jimmy")
	jim.print()

	name := "Bill"
	fmt.Println(*&name)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
func (p *person) print() {
	fmt.Printf("%+v", p)
}
