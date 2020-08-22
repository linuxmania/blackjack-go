package Deck

import (
	//	"fmt"
	"math/rand"
	"time"
)

var deck [52]Card

func GetShuffledDeck() [52]Card {
	createDeck()

	//	fmt.Println(len(deck))
	//	fmt.Println(deck)

	var shuffledDeck [52]Card
	var k = len(deck)
	var tmpDeck []Card = deck[0:52]

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < k; i++ {
		var j = r1.Intn(len(tmpDeck))
		shuffledDeck[i] = tmpDeck[j]
		tmpDeck[j] = tmpDeck[len(tmpDeck)-1]
		tmpDeck = tmpDeck[:len(tmpDeck)-1]
	}
	return shuffledDeck
}

func createDeck() {

	deck[0] = Card{Name: "Ace", Suit: "Spades", Value: 11}
	deck[1] = Card{Name: "Ace", Suit: "Hearts", Value: 11}
	deck[2] = Card{Name: "Ace", Suit: "Clubs", Value: 11}
	deck[3] = Card{Name: "Ace", Suit: "Diamonds", Value: 11}

	deck[4] = Card{Name: "King", Suit: "Spades", Value: 10}
	deck[5] = Card{Name: "King", Suit: "Hearts", Value: 10}
	deck[6] = Card{Name: "King", Suit: "Clubs", Value: 10}
	deck[7] = Card{Name: "King", Suit: "Diamonds", Value: 10}

	deck[8] = Card{Name: "Queen", Suit: "Spades", Value: 10}
	deck[9] = Card{Name: "Queen", Suit: "Hearts", Value: 10}
	deck[10] = Card{Name: "Queen", Suit: "Clubs", Value: 10}
	deck[11] = Card{Name: "Queen", Suit: "Diamonds", Value: 10}

	deck[12] = Card{Name: "Jack", Suit: "Spades", Value: 10}
	deck[13] = Card{Name: "Jack", Suit: "Hearts", Value: 10}
	deck[14] = Card{Name: "Jack", Suit: "Clubs", Value: 10}
	deck[15] = Card{Name: "Jack", Suit: "Diamonds", Value: 10}

	deck[16] = Card{Name: "Ten", Suit: "Spades", Value: 10}
	deck[17] = Card{Name: "Ten", Suit: "Hearts", Value: 10}
	deck[18] = Card{Name: "Ten", Suit: "Clubs", Value: 10}
	deck[19] = Card{Name: "Ten", Suit: "Diamonds", Value: 10}

	deck[20] = Card{Name: "Nine", Suit: "Spades", Value: 9}
	deck[21] = Card{Name: "Nine", Suit: "Hearts", Value: 9}
	deck[22] = Card{Name: "Nine", Suit: "Clubs", Value: 9}
	deck[23] = Card{Name: "Nine", Suit: "Diamonds", Value: 9}

	deck[24] = Card{Name: "Eight", Suit: "Spades", Value: 8}
	deck[25] = Card{Name: "Eight", Suit: "Hearts", Value: 8}
	deck[26] = Card{Name: "Eight", Suit: "Clubs", Value: 8}
	deck[27] = Card{Name: "Eight", Suit: "Diamonds", Value: 8}

	deck[28] = Card{Name: "Seven", Suit: "Spades", Value: 7}
	deck[29] = Card{Name: "Seven", Suit: "Hearts", Value: 7}
	deck[30] = Card{Name: "Seven", Suit: "Clubs", Value: 7}
	deck[31] = Card{Name: "Seven", Suit: "Diamonds", Value: 7}

	deck[32] = Card{Name: "Six", Suit: "Spades", Value: 6}
	deck[33] = Card{Name: "Six", Suit: "Hearts", Value: 6}
	deck[34] = Card{Name: "Six", Suit: "Clubs", Value: 6}
	deck[35] = Card{Name: "Six", Suit: "Diamonds", Value: 6}

	deck[36] = Card{Name: "Five", Suit: "Spades", Value: 5}
	deck[37] = Card{Name: "Five", Suit: "Hearts", Value: 5}
	deck[38] = Card{Name: "Five", Suit: "Clubs", Value: 5}
	deck[39] = Card{Name: "Five", Suit: "Diamonds", Value: 5}

	deck[40] = Card{Name: "Four", Suit: "Spades", Value: 4}
	deck[41] = Card{Name: "Four", Suit: "Hearts", Value: 4}
	deck[42] = Card{Name: "Four", Suit: "Clubs", Value: 4}
	deck[43] = Card{Name: "Four", Suit: "Diamonds", Value: 4}

	deck[44] = Card{Name: "Three", Suit: "Spades", Value: 3}
	deck[45] = Card{Name: "Three", Suit: "Hearts", Value: 3}
	deck[46] = Card{Name: "Three", Suit: "Clubs", Value: 3}
	deck[47] = Card{Name: "Three", Suit: "Diamonds", Value: 3}

	deck[48] = Card{Name: "Two", Suit: "Spades", Value: 2}
	deck[49] = Card{Name: "Two", Suit: "Hearts", Value: 2}
	deck[50] = Card{Name: "Two", Suit: "Clubs", Value: 2}
	deck[51] = Card{Name: "Two", Suit: "Diamonds", Value: 2}
}
