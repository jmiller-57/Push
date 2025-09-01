package gameplay_test

import (
	"testing"

	"github.com/jmill-57/Push/backend/gameplay/deck"
)

func TestIsNaturalTwo(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		Cards    []deck.Card
		suit     deck.Suit
		expected bool
	}{
		{
			name: "Valid natural deck.Two sequence",
			Cards: []deck.Card{
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
			},
			suit:     deck.Hearts,
			expected: true,
		},
		{
			name: "Valid natural deck.Two sequence",
			Cards: []deck.Card{
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Joker, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
			},
			suit:     deck.Hearts,
			expected: true,
		},
		{
			name: "Invalid sequence with wrong suit",
			Cards: []deck.Card{
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{2}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
			},
			suit:     deck.Hearts,
			expected: false,
		},
		{
			name: "Invalid sequence wild deck.Card not acting as a natural deck.Two",
			Cards: []deck.Card{
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
			},
			suit:     deck.Hearts,
			expected: false,
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act: Call the function
			result := deck.IsNaturalTwo(tc.Cards, tc.suit)

			// Assert: Check the result
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func TestValidateRun(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		run      []deck.Card
		suit     deck.Suit
		length   int8
		expected bool
	}{
		// valid runs of deck.Four
		{
			name: "Valid run of deck.Four with deck.Joker wild",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     deck.Hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of deck.Four with deck.Two wild",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     deck.Hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of deck.Four without wilds",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
			},
			suit:     deck.Hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of deck.Four with natural deck.Two",
			run: []deck.Card{
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
			},
			suit:     deck.Hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of deck.Four with natural deck.Two in second position",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
			},
			suit:     deck.Hearts,
			length:   4,
			expected: true,
		},
		// valid runs of deck.Five
		{
			name: "Valid run of deck.Five with deck.Joker wild",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
			},
			suit:     deck.Hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of deck.Five with wild deck.Two",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
			},
			suit:     deck.Hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of deck.Five without wilds",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
			},
			suit:     deck.Hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of deck.Five with natural deck.Two",
			run: []deck.Card{
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
			},
			suit:     deck.Hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of deck.Five with natural deck.Two in second position",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
			},
			suit:     deck.Hearts,
			length:   5,
			expected: true,
		},
		// valid runs of deck.Six
		{
			name: "Valid run of deck.Six with deck.Joker wild",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
				{Rank: deck.Eight, Suit: deck.Hearts, PossibleValues: []int8{8}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of deck.Six with wild deck.Two",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
				{Rank: deck.Eight, Suit: deck.Hearts, PossibleValues: []int8{8}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of deck.Six without wilds",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
				{Rank: deck.Eight, Suit: deck.Hearts, PossibleValues: []int8{8}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of deck.Five with natural deck.Two",
			run: []deck.Card{
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of deck.Six with natural deck.Two in second position",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of deck.Six with natural deck.Two and one wild",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of deck.Six with natural deck.Two and deck.Two wilds",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of deck.Six with deck.Two wilds",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		// valid runs of deck.Seven
		{
			name: "Valid run of deck.Seven with deck.Joker wild",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
				{Rank: deck.Eight, Suit: deck.Hearts, PossibleValues: []int8{8}},
				{Rank: deck.Nine, Suit: deck.Hearts, PossibleValues: []int8{9}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of deck.Seven with wild deck.Two",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
				{Rank: deck.Eight, Suit: deck.Hearts, PossibleValues: []int8{8}},
				{Rank: deck.Nine, Suit: deck.Hearts, PossibleValues: []int8{9}},
			},
			suit:     deck.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of deck.Six without wilds",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
				{Rank: deck.Eight, Suit: deck.Hearts, PossibleValues: []int8{8}},
				{Rank: deck.Nine, Suit: deck.Hearts, PossibleValues: []int8{9}},
			},
			suit:     deck.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of deck.Six with natural deck.Two",
			run: []deck.Card{
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
				{Rank: deck.Eight, Suit: deck.Hearts, PossibleValues: []int8{8}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of deck.Seven with natural deck.Two in second position",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
			},
			suit:     deck.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of deck.Seven with natural deck.Two and one wild",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
			},
			suit:     deck.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of deck.Seven with natural deck.Two and deck.Two wilds",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
			},
			suit:     deck.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of deck.Seven with deck.Two wilds",
			run: []deck.Card{
				{Rank: deck.Ace, Suit: deck.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
			},
			suit:     deck.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of deck.Seven with deck.Two wilds and natural deck.Two",
			run: []deck.Card{
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Two, Suit: deck.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
				{Rank: deck.Seven, Suit: deck.Hearts, PossibleValues: []int8{7}},
			},
			suit:     deck.Hearts,
			length:   7,
			expected: true,
		},
		// invalid runs
		{
			name: "Invalid run not enough Cards",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Hearts, PossibleValues: []int8{5}},
			},
			suit:     deck.Hearts,
			length:   3,
			expected: false,
		},
		{
			name: "Invalid run wrong suited deck.Card",
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Spades, PossibleValues: []int8{5}},
				{Rank: deck.Six, Suit: deck.Hearts, PossibleValues: []int8{6}},
			},
			suit:     deck.Hearts,
			length:   4,
			expected: false,
		},
		{
			name: "Invalid too many wilds",
			run: []deck.Card{
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Spades, PossibleValues: []int8{5}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     deck.Hearts,
			length:   4,
			expected: false,
		},
		{
			name: "Invalid run of 6 too many wilds",
			run: []deck.Card{
				{Rank: deck.Two, Suit: deck.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Four, Suit: deck.Hearts, PossibleValues: []int8{4}},
				{Rank: deck.Five, Suit: deck.Spades, PossibleValues: []int8{5}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: deck.Eight, Suit: deck.Hearts, PossibleValues: []int8{8}},
			},
			suit:     deck.Hearts,
			length:   6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := deck.ValidateRun(tc.run, tc.suit, tc.length)
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func TestValidateBook(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		book     []deck.Card
		expected bool
	}{
		{
			name: "Valid book with deck.Joker wild",
			book: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Three, Suit: deck.Diamonds, PossibleValues: []int8{3}},
				{Rank: deck.Joker, Suit: deck.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			expected: true,
		},
		{
			name: "Valid book with no wilds",
			book: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Three, Suit: deck.Diamonds, PossibleValues: []int8{3}},
				{Rank: deck.Three, Suit: deck.Clubs, PossibleValues: []int8{3}},
			},
			expected: true,
		},
		{
			name: "Valid book of deck.Four without wilds",
			book: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, PossibleValues: []int8{3}},
				{Rank: deck.Three, Suit: deck.Diamonds, PossibleValues: []int8{3}},
				{Rank: deck.Three, Suit: deck.Clubs, PossibleValues: []int8{3}},
				{Rank: deck.Three, Suit: deck.Spades, PossibleValues: []int8{3}},
			},
			expected: true,
		},
		// invalid books
		{
			name: "Invalid book without wilds",
			book: []deck.Card{
				{Rank: deck.Three, Suit: deck.Diamonds, PossibleValues: []int8{3}},
				{Rank: deck.Three, Suit: deck.Clubs, PossibleValues: []int8{3}},
				{Rank: deck.Four, Suit: deck.Spades, PossibleValues: []int8{4}},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := deck.ValidateBook(tc.book)
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
