package main

import (
	"testing"
)

func TestIsNaturalTwo(t *testing.T) {
	// Define test cases
	testCases := []struct {
			name     string
			cards    []Card
			suit     Suit
			expected bool
	}{
			{
				name: "Valid natural two sequence",
				cards: []Card{
						{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
						{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
						{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				},
				suit:     hearts,
				expected: true,
			},
			{
				name: "Valid natural two sequence",
				cards: []Card{
						{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
						{Rank: joker, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
						{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				},
				suit:     hearts,
				expected: true,
			},
			{
				name: "Invalid sequence with wrong suit",
				cards: []Card{
						{Rank: two, Suit: spades, PossibleValues: []int8{2}},
						{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
						{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				},
				suit:     hearts,
				expected: false,
			},
			{
				name: "Invalid sequence wild card not acting as a natural two",
				cards: []Card{
						{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
						{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
						{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				},
				suit:     hearts,
				expected: false,
			},
	}

	// Iterate over test cases
	for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
					// Act: Call the function
					result := IsNaturalTwo(tc.cards, tc.suit)

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
		run      []Card
		suit     Suit
		length   int8
		expected bool
	}{
		// valid runs of four
		{
			name: "Valid run of four with joker wild",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of four with two wild",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: two, Suit: spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of four without wilds",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
			},
			suit:     hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of four with natural two",
			run: []Card{
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
			},
			suit:     hearts,
			length:   4,
			expected: true,
		},
		{
			name: "Valid run of four with natural two in second position",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
			},
			suit:     hearts,
			length:   4,
			expected: true,
		},
		// valid runs of five
		{
			name: "Valid run of five with joker wild",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
			},
			suit:     hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of five with wild two",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: two, Suit: spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
			},
			suit:     hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of five without wilds",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
			},
			suit:     hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of five with natural two",
			run: []Card{
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
			},
			suit:     hearts,
			length:   5,
			expected: true,
		},
		{
			name: "Valid run of five with natural two in second position",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
			},
			suit:     hearts,
			length:   5,
			expected: true,
		},
		// valid runs of six
		{
			name: "Valid run of six with joker wild",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
				{Rank: eight, Suit: hearts, PossibleValues: []int8{8}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of six with wild two",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: two, Suit: spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
				{Rank: eight, Suit: hearts, PossibleValues: []int8{8}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of six without wilds",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
				{Rank: eight, Suit: hearts, PossibleValues: []int8{8}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of five with natural two",
			run: []Card{
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of six with natural two in second position",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of six with natural two and one wild",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of six with natural two and two wilds",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of six with two wilds",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		// valid runs of seven
		{
			name: "Valid run of seven with joker wild",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
				{Rank: eight, Suit: hearts, PossibleValues: []int8{8}},
				{Rank: nine, Suit: hearts, PossibleValues: []int8{9}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of seven with wild two",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: two, Suit: spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
				{Rank: eight, Suit: hearts, PossibleValues: []int8{8}},
				{Rank: nine, Suit: hearts, PossibleValues: []int8{9}},
			},
			suit:     hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of six without wilds",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
				{Rank: eight, Suit: hearts, PossibleValues: []int8{8}},
				{Rank: nine, Suit: hearts, PossibleValues: []int8{9}},
			},
			suit:     hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of six with natural two",
			run: []Card{
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
				{Rank: eight, Suit: hearts, PossibleValues: []int8{8}},
			},
			suit:     hearts,
			length:   6,
			expected: true,
		},
		{
			name: "Valid run of seven with natural two in second position",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
			},
			suit:     hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of seven with natural two and one wild",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
			},
			suit:     hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of seven with natural two and two wilds",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
			},
			suit:     hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of seven with two wilds",
			run: []Card{
				{Rank: ace, Suit: hearts, PossibleValues: []int8{1, 14}},
				{Rank: two, Suit: spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
			},
			suit:     hearts,
			length:   7,
			expected: true,
		},
		{
			name: "Valid run of seven with two wilds and natural two",
			run: []Card{
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: two, Suit: hearts, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: two, Suit: spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
				{Rank: seven, Suit: hearts, PossibleValues: []int8{7}},
			},
			suit:     hearts,
			length:   7,
			expected: true,
		},
		// invalid runs
		{
			name: "Invalid run not enough cards",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: hearts, PossibleValues: []int8{5}},
			},
			suit:     hearts,
			length:   3,
			expected: false,
		},
		{
			name: "Invalid run wrong suited card",
			run: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: spades, PossibleValues: []int8{5}},
				{Rank: six, Suit: hearts, PossibleValues: []int8{6}},
			},
			suit:     hearts,
			length:   4,
			expected: false,
		},
		{
			name: "Invalid too many wilds",
			run: []Card{
				{Rank: two, Suit: spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: spades, PossibleValues: []int8{5}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			suit:     hearts,
			length:   4,
			expected: false,
		},
		{
			name: "Invalid run of 6 too many wilds",
			run: []Card{
				{Rank: two, Suit: spades, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: four, Suit: hearts, PossibleValues: []int8{4}},
				{Rank: five, Suit: spades, PossibleValues: []int8{5}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: eight, Suit: hearts, PossibleValues: []int8{8}},
			},
			suit:     hearts,
			length:   6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidateRun(tc.run, tc.suit, tc.length)
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
		book     []Card
		expected bool
	}{
		{
			name: "Valid book with joker wild",
			book: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: three, Suit: diamonds, PossibleValues: []int8{3}},
				{Rank: joker, Suit: anySuit, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			expected: true,
		},
		{
			name: "Valid book with no wilds",
			book: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: three, Suit: diamonds, PossibleValues: []int8{3}},
				{Rank: three, Suit: clubs, PossibleValues: []int8{3}},
			},
			expected: true,
		},
		{
			name: "Valid book of four without wilds",
			book: []Card{
				{Rank: three, Suit: hearts, PossibleValues: []int8{3}},
				{Rank: three, Suit: diamonds, PossibleValues: []int8{3}},
				{Rank: three, Suit: clubs, PossibleValues: []int8{3}},
				{Rank: three, Suit: spades, PossibleValues: []int8{3}},
			},
			expected: true,
		},
		// invalid books
		{
			name: "Invalid book without wilds",
			book: []Card{
				{Rank: three, Suit: diamonds, PossibleValues: []int8{3}},
				{Rank: three, Suit: clubs, PossibleValues: []int8{3}},
				{Rank: four, Suit: spades, PossibleValues: []int8{4}},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing	.T) {
			result := ValidateBook(tc.book)
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}