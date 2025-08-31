package gameplay

import "slices"

type SetType string

const (
	Book SetType = "book"
	Run  SetType = "run"
)

type SetRequirement struct {
	Type      SetType
	MinLength int8
}

type Play struct {
	Cards []Card
	Set   SetRequirement
}

func FindBooks(hand []Card) [][]Card {
	rankGroups := make(map[Rank][]Card)
	var wildcards []Card

	for _, card := range hand {
		if card.IsWild() {
			wildcards = append(wildcards, card)
		} else {
			rankGroups[card.Rank] = append(rankGroups[card.Rank], card)
		}
	}

	var books [][]Card
	usedWilds := 0

	for _, group := range rankGroups {
		n := len(group)
		if n >= 3 { // pure book
			books = append(books, group)
		} else if n == 2 && usedWilds < len(wildcards) {
			book := append([]Card{}, group...)
			book = append(book, wildcards[usedWilds])
			usedWilds++
			books = append(books, book)
		}
	}
	return books
}

func FindRunsWithWilds(hand []Card) [][]Card {
	suitGroups := make(map[Suit][]Card)
	var wilds []Card
	var runs [][]Card

	// Separate wildcards and group the rest by suit
	for _, card := range hand {
		if card.IsWild() && !(card.Rank == Two) {
			wilds = append(wilds, card)
		} else {
			suitGroups[card.Suit] = append(suitGroups[card.Suit], card)
		}
	}

	// Helper: is a card being used naturally (not counting as a wildcard)?
	isNatural := func(card Card, suit Suit, value int8) bool {
		if !card.IsWild() {
			return true
		}
		if card.Rank == Two && card.Suit == suit && value == 2 {
			return true // natural two
		}
		return false
	}

	// Attempt runs in each suit
	for suit, cards := range suitGroups {
		// Build value → []Card lookup map
		valToCards := make(map[int8][]Card)
		for _, card := range cards {
			for _, val := range card.PossibleValues {
				valToCards[val] = append(valToCards[val], card)
			}
		}

		// Try all run lengths from 4 to 7
		for start := 1; start <= 11; start++ {
			for length := 4; length <= 7; length++ {
				end := start + length - 1
				if end > 14 {
					continue // no run past Ace high
				}

				var run []Card
				usedCardIDs := map[string]bool{}
				wildCount := 0
				wildIdx := 0

				for val := int8(start); val <= int8(end); val++ {
					var usedCard *Card
					for _, cand := range valToCards[val] {
						id := cand.String()
						if !usedCardIDs[id] {
							usedCard = &cand
							break
						}
					}

					if usedCard != nil {
						run = append(run, *usedCard)
						usedCardIDs[usedCard.String()] = true

						if !isNatural(*usedCard, suit, val) {
							wildCount++
						}
					} else if wildIdx < len(wilds) {
						run = append(run, wilds[wildIdx])
						wildCount++
						wildIdx++
					} else {
						break
					}
				}

				if len(run) == length {
					if wildCount <= 1 || (length >= 6 && wildCount == 2) {
						runs = append(runs, run)
					}
				}
			}
		}
	}

	return runs
}

func ValidateBook(book []Card) bool {
	if len(book) < 3 {
		return false
	}

	var rank Rank

	if !book[0].IsWild() {
		rank = book[0].Rank
	} else if !book[1].IsWild() {
		rank = book[1].Rank
	} else {
		// cannot have more than 1 wild in a book
		return false
	}

	numWilds := 0

	for _, card := range book {
		if numWilds > 1 || (card.IsWild() && numWilds > 0) {
			return false // cannot have more than 1 wild in a book
		}
		if card.IsWild() {
			numWilds++
			continue
		}
		if card.Rank != rank {
			return false
		}
	}
	return true
}

func ValidateRun(run []Card, suit Suit, length int8) bool {
	if len(run) < 4 {
		return false
	}

	wildsUsed := 0
	cardsAssesed := 0

	for _, card := range run {
		if !card.IsWild() && card.Suit != suit {
			return false
		}
		if card.IsWild() {
			wildsUsed++
		}

		// run of length 4: 0, 1, 2,

		if cardsAssesed < (len(run) - 1) {
			if cardsAssesed < 2 && run[cardsAssesed].Rank == Two && IsNaturalTwo(run[cardsAssesed:], suit) {
				wildsUsed--
			}

			next := run[cardsAssesed+1]

			found := false
			for _, curr := range card.PossibleValues {
				if slices.Contains(next.PossibleValues, curr+1) {
					// valid run sequence
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		cardsAssesed++
	}

	return true
}

func IsNaturalTwo(cards []Card, suit Suit) bool {
	return cards[0].Rank == Two && cards[0].Suit == suit &&
		((cards[1].Rank == Three && cards[1].Suit == suit) ||
			(cards[1].IsWild() && cards[2].Rank == Four && cards[2].Suit == suit))
}
