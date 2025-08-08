package main

type Suit string
type Rank string

const (
	hearts   Suit = "♥"
	spades   Suit = "♠"
	diamonds Suit = "♦"
	clubs    Suit = "♣"
	anySuit  Suit = "*"
)

const (
	ace   Rank = "A"
	two   Rank = "2"
	three Rank = "3"
	four  Rank = "4"
	five  Rank = "5"
	six   Rank = "6"
	seven Rank = "7"
	eight Rank = "8"
	nine  Rank = "9"
	ten   Rank = "10"
	jack  Rank = "J"
	queen Rank = "Q"
	king  Rank = "K"
	joker Rank = "Joker"
)

type Card struct {
	Rank  Rank
	Suit  Suit
	Points int8
	PossibleValues []int8
}

func NewCard(suit Suit, rank Rank) Card {
	return Card{
		Suit:  suit,
		Rank:  rank,
		Points: pointsFromRank(rank),
		PossibleValues: valuesFromRank(rank),
	}
}

func pointsFromRank(rank Rank) int8 {
	switch rank {
	case three, four, five, six, seven, eight, nine:
		return 5
	case ace, ten, jack, queen, king:
		return 10
	case two:
		return 20
	case joker:
		return 50
	default:
		return 0
	}
}

func valuesFromRank(rank Rank) []int8 {
	switch rank {
	case ace:
		return []int8{1, 14}
	case three:
		return []int8{3}
	case four:
		return []int8{4}
	case five:
		return []int8{5}
	case six:
		return []int8{6}
	case seven:
		return []int8{7}
	case eight:
		return []int8{8}
	case nine:
		return []int8{9}
	case ten:
		return []int8{10}
	case jack:
		return []int8{11}
	case queen:
		return []int8{12}
	case king:
		return []int8{13}
	case joker, two:
		return []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	// wildcard, so it has all values
	}
	return []int8{}
}

func (c Card) IsWild() bool {
	return c.Rank == joker || c.Rank == two
}

func (c Card) String() string {
	return string(c.Rank) + string(c.Suit)
}

func (c Card) Equals(other Card) bool {
	return c.Rank == other.Rank && c.Suit == other.Suit
}
