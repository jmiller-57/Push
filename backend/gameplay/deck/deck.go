package deck

import (
	"math/rand"
	"time"
	//"fmt"
)

type Deck []Card

func NewDeck() Deck {
	suits := []Suit{Hearts, Diamonds, Spades, Clubs}
	ranks := []Rank{
		Ace, Two, Three, Four, Five, Six, Seven,
		Eight, Nine, Ten, Jack, Queen, King, Joker,
	}

	var deck Deck

	for _, suit := range suits {
		for _, rank := range ranks {
			if rank != Joker {
				deck = append(deck, NewCard(suit, rank))
				deck = append(deck, NewCard(suit, rank))
			} else {
				//fmt.Printf("joker: rank=%v, suit=%v", rank, suit)
				deck = append(deck, NewCard(AnySuit, rank))
			}
		}
	}

	deck.ShuffleDeck()
	return deck
}

func (d *Deck) ShuffleDeck() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for range 7 {
		r.Shuffle(len(*d), func(i, j int) {
			(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
		})
	}
}

func (d *Deck) DrawCard() Card {
	if len(*d) == 0 {
		return Card{} // or handle empty deck case
	}
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}
