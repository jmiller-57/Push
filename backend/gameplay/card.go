package gameplay

type Suit string
type Rank string

const (
	Hearts   Suit = "♥"
	Spades   Suit = "♠"
	Diamonds Suit = "♦"
	Clubs    Suit = "♣"
	AnySuit  Suit = "*"
)

const (
	Ace   Rank = "A"
	Two   Rank = "2"
	Three Rank = "3"
	Four  Rank = "4"
	Five  Rank = "5"
	Six   Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "J"
	Queen Rank = "Q"
	King  Rank = "K"
	Joker Rank = "Joker"
)

type Card struct {
	Rank           Rank
	Suit           Suit
	Points         int8
	PossibleValues []int8
}

func NewCard(suit Suit, rank Rank) Card {
	return Card{
		Suit:           suit,
		Rank:           rank,
		Points:         pointsFromRank(rank),
		PossibleValues: valuesFromRank(rank),
	}
}

func pointsFromRank(rank Rank) int8 {
	switch rank {
	case Three, Four, Five, Six, Seven, Eight, Nine:
		return 5
	case Ace, Ten, Jack, Queen, King:
		return 10
	case Two:
		return 20
	case Joker:
		return 50
	default:
		return 0
	}
}

func valuesFromRank(rank Rank) []int8 {
	switch rank {
	case Ace:
		return []int8{1, 14}
	case Three:
		return []int8{3}
	case Four:
		return []int8{4}
	case Five:
		return []int8{5}
	case Six:
		return []int8{6}
	case Seven:
		return []int8{7}
	case Eight:
		return []int8{8}
	case Nine:
		return []int8{9}
	case Ten:
		return []int8{10}
	case Jack:
		return []int8{11}
	case Queen:
		return []int8{12}
	case King:
		return []int8{13}
	case Joker, Two:
		return []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
		// wildcard, so it has all values
	}
	return []int8{}
}

func (c Card) IsWild() bool {
	return c.Rank == Joker || c.Rank == Two
}

func (c Card) String() string {
	return string(c.Rank) + string(c.Suit)
}

func (c Card) Equals(other Card) bool {
	return c.Rank == other.Rank && c.Suit == other.Suit
}
