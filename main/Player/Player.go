package Player

import (
	"bufio"
	"fmt"
	"main/Deck"
	"os"
)

type Player struct {
	Name     string
	numCards int
	Hand     [10]Deck.Card
	IsBust   bool
}

func (p Player) ReportResults(isDealerBust bool, dealerValue int) {
	fmt.Println(p.Name)
	if p.GetValue() > 21 {
		fmt.Println("Busted!")
	} else if isDealerBust == true {
		fmt.Println("Winner - Dealer Busted!!")
	} else if p.GetValue() > dealerValue {
		fmt.Println("Winner !!")
	} else if p.GetValue() < dealerValue {
		fmt.Println("Loser :-(")
	} else {
		fmt.Println("Tie - Push :-|")
	}
	fmt.Println()
}

func Process(p1 *Player, d *[52]Deck.Card, pos *int) {
	for {
		var a = ""
		p1.ReportHand()
		fmt.Println()
		fmt.Println("hit?")
		reader := bufio.NewReader(os.Stdin)
		a, _ = reader.ReadString('\n')
		if a == "y\n" {
			AddCard(d[*pos], p1)
			*pos += 1
			if p1.GetValue() > 21 {
				if p1.devalueAce(p1.Hand[0:10]) == false {
					p1.IsBust = true
					fmt.Println()
					p1.ReportHand()
					fmt.Println("Busted!")
					fmt.Println()
					return
				}
			}
		} else {
			return
		}
	} // infinite loop either bust or opt out
}

func (p Player) devalueAce(tmpHand []Deck.Card) bool {
	for i := 0; i < p.numCards; i++ {
		if tmpHand[i].Name == "Ace" && tmpHand[i].Value == 11 {
			tmpHand[i].Value = 1
			return true
		}
	}
	return false
}

func (p Player) ReportHand() {
	fmt.Println(p.Name)
	for i := 0; i < p.numCards; i++ {
		p.Hand[i].PrintCard()
	}
	fmt.Println("Value", p.GetValue())
}

func (p Player) GetValue() int {
	var value = 0
	for i := 0; i < p.numCards; i++ {
		value += p.Hand[i].Value
	}
	return value
}

func AddCard(c Deck.Card, p *Player) {
	p.Hand[p.numCards] = c
	p.numCards += 1
}
