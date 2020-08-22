package main

import (
	"fmt"
	"main/Deck"
	"main/Player"
)

func main() {
	var deck = Deck.GetShuffledDeck()
	var position = 0

	//	fmt.Println(len(deck))
	//	fmt.Println(deck)

	p1 := Player.Player{Name: "Player1"}
	p2 := Player.Player{Name: "Player2"}
	dp := Player.Player{Name: "Dealer"}
	d := Player.Dealer{P: dp}

	Player.AddCard(deck[position], &p1)
	position += 1
	Player.AddCard(deck[position], &p2)
	position += 1
	Player.AddCard(deck[position], &d.P)
	position += 1

	Player.AddCard(deck[position], &p1)
	position += 1
	Player.AddCard(deck[position], &p2)
	position += 1
	Player.AddCard(deck[position], &d.P)
	position += 1

	p1.ReportHand()
	fmt.Println()
	p2.ReportHand()
	fmt.Println()
	d.ReportShowing()
	fmt.Println()

	Player.Process(&p1, &deck, &position)
	Player.Process(&p2, &deck, &position)
	Player.Process(&d.P, &deck, &position)

	p1.ReportResults(d.P.IsBust, d.P.GetValue())
	p2.ReportResults(d.P.IsBust, d.P.GetValue())

}
