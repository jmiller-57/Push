package gameplay_test

import (
	"testing"

	"github.com/jmiller-57/Push/backend/gameplay"
)

var g gameplay.GameState

func TestStartGame(t *testing.T) {
	playerNames := []string{"Alice", "Bob", "Charlie"}
	game := g.NewGame(playerNames)

	// Assert: Check initial game state
	if len(game.Players) != len(playerNames) {
		t.Errorf("Expected %d players, got %d", len(playerNames), len(game.Players))
	}

	if len(game.Players[0].Hand) != 11 {
		t.Errorf("Expected 11 cards in each player's hand, got %d", len(game.Players[0].Hand))
	}

	if game.FaceUpCard == nil {
		t.Errorf("Expected a face-up card, got nil")
	}

	if !game.HasStarted {
		t.Errorf("Expected game to be started, got HasStarted = false")
	}
}

func TestTakeFaceUpCard(t *testing.T) {
	playerNames := []string{"Alice", "Bob"}
	game := g.NewGame(playerNames)

	// Act: Current player takes the face-up card
	game.TakeFaceUpCard()

	// Assert: Check the player's hand and face-up card
	if len(game.Players[0].Hand) != 12 {
		t.Errorf("Expected 12 cards in player's hand, got %d", len(game.Players[0].Hand))
	}

	if game.FaceUpCard != nil {
		t.Errorf("Expected face-up card to be nil, got %v", game.FaceUpCard)
	}
}

func TestPushFaceUpCard(t *testing.T) {
	playerNames := []string{"Alice", "Bob"}
	game := g.NewGame(playerNames)

	// Act: Current player pushes the face-up card
	game.PushFaceUpCard()

	// Assert: Check the next player's hand and current player's hand
	if len(game.Players[1].Hand) != 13 {
		t.Errorf("Expected 13 cards in next player's hand, got %d", len(game.Players[1].Hand))
	}

	if len(game.Players[0].Hand) != 12 {
		t.Errorf("Expected 12 cards in current player's hand, got %d", len(game.Players[0].Hand))
	}

	if game.FaceUpCard != nil {
		t.Errorf("Expected face-up card to be nil, got %v", game.FaceUpCard)
	}
}

func TestDiscard(t *testing.T) {
	playerNames := []string{"Alice", "Bob"}
	game := g.NewGame(playerNames)
	game.FaceUpCard = nil

	// Arrange: Get a card from the current player's hand
	cardToDiscard := game.Players[0].Hand[0]

	// Act: Current player discards the card
	game.Discard(cardToDiscard)

	// Assert: Check the player's hand and face-up card
	if len(game.Players[0].Hand) != 10 {
		t.Errorf("Expected 10 cards in player's hand, got %d", len(game.Players[0].Hand))
	}

	if game.FaceUpCard == nil || !game.FaceUpCard.Equals(cardToDiscard) {
		t.Errorf("Expected face-up card to be %v, got %v", cardToDiscard, game.FaceUpCard)
	}
}

func TestPlayRun(t *testing.T) {
	playerNames := []string{"Alice", "Bob"}
	game := g.NewGame(playerNames)

	testCase := []struct {
		name        string
		hand        []gameplay.Card // The player's hand
		run         []gameplay.Card // The run to play
		expectValid bool            // Whether the run is expected to be valid
	}{
		{
			name: "Valid run of gameplay.Four",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{6}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{9}},
				{Rank: gameplay.Ten, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{10}},
			},
			run: []gameplay.Card{
				{Rank: gameplay.Four, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{6}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			expectValid: true,
		},
		{
			name: "Valid run of gameplay.Five",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{6}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{9}},
				{Rank: gameplay.Ten, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{10}},
				{Rank: gameplay.Jack, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{11}},
			},
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{6}},
				{Rank: gameplay.Seven, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{7}},
			},
			expectValid: true,
		},
		{
			name: "Valid run of gameplay.Six",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Clubs, Points: 5, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{6}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{9}},
				{Rank: gameplay.Ten, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{10}},
				{Rank: gameplay.Jack, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{11}},
				{Rank: gameplay.Queen, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{12}},
			},
			run: []gameplay.Card{
				{Rank: gameplay.Seven, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{9}},
				{Rank: gameplay.Ten, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{10}},
				{Rank: gameplay.Jack, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{11}},
				{Rank: gameplay.Queen, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{12}},
			},
			expectValid: true,
		},
		{
			name: "Valid run of gameplay.Seven",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Clubs, Points: 5, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{6}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{9}},
				{Rank: gameplay.Ten, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{10}},
				{Rank: gameplay.Jack, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{11}},
				{Rank: gameplay.Queen, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{12}},
				{Rank: gameplay.King, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{13}},
			},
			run: []gameplay.Card{
				{Rank: gameplay.Seven, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{9}},
				{Rank: gameplay.Ten, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{10}},
				{Rank: gameplay.Jack, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{11}},
				{Rank: gameplay.Queen, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{12}},
				{Rank: gameplay.King, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{13}},
			},
			expectValid: true,
		},
		{
			name: "Invalid run with non-consecutive ranks",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{6}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{9}},
				{Rank: gameplay.Ten, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{10}},
			},
			run: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{3}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{6}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
			},
			expectValid: false,
		},
		{
			name: "Invalid run not long enough",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{3}},
				{Rank: gameplay.Four, Suit: gameplay.Clubs, Points: 5, PossibleValues: []int8{4}},
				{Rank: gameplay.Five, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{5}},
				{Rank: gameplay.Six, Suit: gameplay.Hearts, Points: 5, PossibleValues: []int8{6}},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				{Rank: gameplay.Seven, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{9}},
				{Rank: gameplay.Ten, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{10}},
			},
			run: []gameplay.Card{
				{Rank: gameplay.Seven, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{7}},
				{Rank: gameplay.Eight, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{8}},
				{Rank: gameplay.Nine, Suit: gameplay.Spades, Points: 5, PossibleValues: []int8{9}},
				{Rank: gameplay.Ten, Suit: gameplay.Spades, Points: 10, PossibleValues: []int8{10}},
			},
			expectValid: false,
		},
	}

	var roundIdx int8 = 3

	// Iterate over test cases
	for _, tc := range testCase {
		if roundIdx == 6 {
			roundIdx = 3 // Reset round index for next test case
		}
		t.Run(tc.name, func(t *testing.T) {
			// Arrange: Set up the player's hand
			game = g.NewGame(playerNames)
			game.CurrentTurnIdx = 0 // Ensure we are testing the first player
			game.Players[game.CurrentTurnIdx].Hand = tc.hand
			tc.hand = append(tc.hand, *game.FaceUpCard) // Add the face-up card to the hand

			game.TakeFaceUpCard() // Ensure the player has a face-up card

			suit := tc.run[0].Suit
			roundObj := gameplay.RoundObjective{
				RoundNumber: roundIdx + 1,
				Sets:        gameplay.RoundObjectives[roundIdx].Sets,
			}
			game.Round = roundObj
			runLength := gameplay.RoundObjectives[roundIdx].Sets[gameplay.RUN_INDEX].MinLength

			// Act: Attempt to play the run
			if tc.expectValid {

				game.PlayRun(tc.run, suit, runLength)

				// Assert: Check the player's hand and downed sets
				if len(game.Players[0].Hand) != len(tc.hand)-len(tc.run) {
					t.Errorf("Expected %d cards in player's hand after playing run, got %d",
						len(tc.hand)-len(tc.run), len(game.Players[0].Hand))
				}

				if len(game.DownedSets[game.Players[0].ID]) != 1 {
					t.Errorf("Expected 1 downed set for player, got %d", len(game.DownedSets[game.Players[0].ID]))
				}
			} else {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for invalid run, but no panic occurred")
					}
				}()
				game.PlayRun(tc.run, suit, runLength)
			}
		})
		roundIdx++
	}
}

func TestPlayBook(t *testing.T) {
	playerNames := []string{"Alice", "Bob"}
	game := g.NewGame(playerNames)

	// Define test cases
	testCases := []struct {
		name        string
		hand        []gameplay.Card // The player's hand
		book        []gameplay.Card // The book to play
		expectValid bool            // Whether the book is expected to be valid
	}{
		{
			name: "Valid book with gameplay.Three cards of the same rank",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Three, Suit: gameplay.Spades},
				{Rank: gameplay.Three, Suit: gameplay.Diamonds},
				{Rank: gameplay.Four, Suit: gameplay.Clubs},
			},
			book: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Three, Suit: gameplay.Spades},
				{Rank: gameplay.Three, Suit: gameplay.Diamonds},
			},
			expectValid: true,
		},
		{
			name: "Invalid book with mixed ranks",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Four, Suit: gameplay.Spades},
				{Rank: gameplay.Five, Suit: gameplay.Diamonds},
				{Rank: gameplay.Six, Suit: gameplay.Clubs},
			},
			book: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Four, Suit: gameplay.Spades},
				{Rank: gameplay.Five, Suit: gameplay.Diamonds},
			},
			expectValid: false,
		},
		{
			name: "Invalid book with less than gameplay.Three cards",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Three, Suit: gameplay.Spades},
				{Rank: gameplay.Four, Suit: gameplay.Diamonds},
			},
			book: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Three, Suit: gameplay.Spades},
			},
			expectValid: false,
		},
		{
			name: "Valid book with one wildcard",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Three, Suit: gameplay.Spades},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit},
				{Rank: gameplay.Four, Suit: gameplay.Clubs},
			},
			book: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Three, Suit: gameplay.Spades},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit},
			},
			expectValid: true,
		},
		{
			name: "Invalid book with gameplay.Two wildcards",
			hand: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit},
				{Rank: gameplay.Four, Suit: gameplay.Clubs},
			},
			book: []gameplay.Card{
				{Rank: gameplay.Three, Suit: gameplay.Hearts},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit},
				{Rank: gameplay.Joker, Suit: gameplay.AnySuit},
			},
			expectValid: false,
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange: Set up the player's hand
			game = g.NewGame(playerNames)
			game.CurrentTurnIdx = 0 // Ensure we are testing the first player
			game.Players[game.CurrentTurnIdx].Hand = tc.hand
			tc.hand = append(tc.hand, *game.FaceUpCard) // Add the face-up card to the hand

			game.TakeFaceUpCard() // Ensure the player has a face-up card

			// Act: Attempt to play the book
			if tc.expectValid {
				game.PlayBook(tc.book)

				// Assert: Check the player's hand and downed sets
				if len(game.Players[0].Hand) != len(tc.hand)-len(tc.book) {
					t.Errorf("Expected %d cards in player's hand after playing book, got %d",
						len(tc.hand)-len(tc.book), len(game.Players[0].Hand))
				}

				if len(game.DownedSets[game.Players[0].ID]) != 1 {
					t.Errorf("Expected 1 downed set for player, got %d", len(game.DownedSets[game.Players[0].ID]))
				}
			} else {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for invalid book, but no panic occurred")
					}
				}()
				game.PlayBook(tc.book)
			}
		})
	}
}
