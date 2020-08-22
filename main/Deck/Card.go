package Deck

import (
	"fmt"
)

type Card struct {
	Name  string
	Suit  string
	Value int
}

func (c Card) PrintCard() {
	fmt.Println(c.Name, "of", c.Suit)
}
