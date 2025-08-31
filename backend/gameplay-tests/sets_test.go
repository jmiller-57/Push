package gameplay_test

import (
	"testing"
	"github.com/jmiller-57/Push/push-backend/gameplay"
)

func TestIsNaturalTwo(t *testing.T) {
	// Define test cases
	testCases := []struct {
			name     string
			Cards    []gameplay.Card
			suit     gameplay.Suit
			expected bool
	}{
			{
				name: "Valid natural gameplay.Two sequence",
				Cards: []gameplay.Card{
						{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
						{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
						{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				},
				suit:     gameplay.Hearts,
				expected: true,
			},
			{
				name: "Valid natural gameplay.Two sequence",
				Cards: []gameplay.Card{
						{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
						{Rank: gameplay.Joker, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
						{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				},
				suit:     gameplay.Hearts,
				expected: true,
			},
			{
				name: "Invalid sequence with wrong suit",
				Cards: []gameplay.Card{
						{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{2}},
						{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
						{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				},
				suit:     gameplay.Hearts,
				expected: false,
			},
			{
				name: "Invalid sequence wild gameplay.Card not acting as a natural gameplay.Two",
				Cards: []gameplay.Card{
						{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
						{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
						{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				},
				suit:     gameplay.Hearts,
				expected: false,
			},
	}

	// Iterate over test cases
	for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
					// Act: Call the function
					result := gameplay.IsNaturalTwo(tc.Cards, tc.suit)

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
		run      []gameplay.Card
		suit     gameplay.Suit
		length   int8
		expected bool
	}{
		// valid runs of gameplay.Four
		{
			name: "Valid run of gameplay.Four with gameplay.Joker wild",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     gameplay.Hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Four with gameplay.Two wild",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     gameplay.Hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Four without wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
			},
			suit:     gameplay.Hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Four with natural gameplay.Two",
			run: []gameplay.Card{
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
			},
			suit:     gameplay.Hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Four with natural gameplay.Two in second position",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
			},
			suit:     gameplay.Hearts,
			length:   4,
			expected: true,
		},
		// valid runs of gameplay.Five
		{
			name: "Valid run of gameplay.Five with gameplay.Joker wild",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
			},
			suit:     gameplay.Hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Five with wild gameplay.Two",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
			},
			suit:     gameplay.Hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Five without wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
			},
			suit:     gameplay.Hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Five with natural gameplay.Two",
			run: []gameplay.Card{
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
			},
			suit:     gameplay.Hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Five with natural gameplay.Two in second position",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
			},
			suit:     gameplay.Hearts,
			length:   5,
			expected: true,
		},
		// valid runs of gameplay.Six
		{
			name: "Valid run of gameplay.Six with gameplay.Joker wild",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Hearts, PossibleValues: []int8{8}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Six with wild gameplay.Two",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Hearts, PossibleValues: []int8{8}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Six without wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Hearts, PossibleValues: []int8{8}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Five with natural gameplay.Two",
			run: []gameplay.Card{
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Six with natural gameplay.Two in second position",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Six with natural gameplay.Two and one wild",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Six with natural gameplay.Two and gameplay.Two wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Six with gameplay.Two wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		// valid runs of gameplay.Seven
		{
			name: "Valid run of gameplay.Seven with gameplay.Joker wild",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Hearts, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Hearts, PossibleValues: []int8{9}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Seven with wild gameplay.Two",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Hearts, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Hearts, PossibleValues: []int8{9}},
			},
			suit:     gameplay.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Six without wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Hearts, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Hearts, PossibleValues: []int8{9}},
			},
			suit:     gameplay.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Six with natural gameplay.Two",
			run: []gameplay.Card{
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Hearts, PossibleValues: []int8{8}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Seven with natural gameplay.Two in second position",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
			},
			suit:     gameplay.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Seven with natural gameplay.Two and one wild",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
			},
			suit:     gameplay.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Seven with natural gameplay.Two and gameplay.Two wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
			},
			suit:     gameplay.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Seven with gameplay.Two wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Ace, Suit: gameplay.Hearts, PossibleValues: []int8{1, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
			},
			suit:     gameplay.Hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of gameplay.Seven with gameplay.Two wilds and natural gameplay.Two",
			run: []gameplay.Card{
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, PossibleValues: []int8{7}},
			},
			suit:     gameplay.Hearts,
			length:   7,
			expected: true,
		},
		// invalid runs
		{
			name: "Invalid run not enough Cards",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, PossibleValues: []int8{5}},
			},
			suit:     gameplay.Hearts,
			length:   3,
			expected: false,
		},
		{
			name: "Invalid run wrong suited gameplay.Card",
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Spades, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, PossibleValues: []int8{6}},
			},
			suit:     gameplay.Hearts,
			length:   4,
			expected: false,
		},
		{
			name: "Invalid too many wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Spades, PossibleValues: []int8{5}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     gameplay.Hearts,
			length:   4,
			expected: false,
		},
		{
			name: "Invalid run of 6 too many wilds",
			run: []gameplay.Card{
				{Rank: gameplay.Two, Suit: gameplay.Spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Spades, PossibleValues: []int8{5}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Eight, Suit: gameplay.Hearts, PossibleValues: []int8{8}},
			},
			suit:     gameplay.Hearts,
			length:   6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := gameplay.ValidateRun(tc.run, tc.suit, tc.length)
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
		book     []gameplay.Card
		expected bool
	}{
		{
			name: "Valid book with gameplay.Joker wild",
			book: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Three, Suit: gameplay.Diamonds, PossibleValues: []int8{3}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			expected: true,
		},
		{
			name: "Valid book with no wilds",
			book: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Three, Suit: gameplay.Diamonds, PossibleValues: []int8{3}},
				{Rank: gameplay.Three, Suit: gameplay.Clubs, PossibleValues: []int8{3}},
			},
			expected: true,
		},
		{
			name: "Valid book of gameplay.Four without wilds",
			book: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, PossibleValues: []int8{3}},
				{Rank: gameplay.Three, Suit: gameplay.Diamonds, PossibleValues: []int8{3}},
				{Rank: gameplay.Three, Suit: gameplay.Clubs, PossibleValues: []int8{3}},
				{Rank: gameplay.Three, Suit: gameplay.Spades, PossibleValues: []int8{3}},
			},
			expected: true,
		},
		// invalid books
		{
			name: "Invalid book without wilds",
			book: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Diamonds, PossibleValues: []int8{3}},
				{Rank: gameplay.Three, Suit: gameplay.Clubs, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Spades, PossibleValues: []int8{4}},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing	.T) {
			result := gameplay.ValidateBook(tc.book)
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}