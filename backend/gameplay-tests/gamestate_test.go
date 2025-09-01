package gameplay_test

import (
	"testing"

	"github.com/jmiller-57/Push/backend/gameplay"
	"github.com/jmiller-57/Push/backend/gameplay/deck"
)

var g gameplay.GameState

func TestStartGame(t *testing.T) {
	playerNames := []string{"Alice", "Bob", "Charlie"}
	game := gameplay.NewGame(playerNames)

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
	game := gameplay.NewGame(playerNames)

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
	game := gameplay.NewGame(playerNames)

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
	game := gameplay.NewGame(playerNames)
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
	game := gameplay.NewGame(playerNames)

	testCase := []struct {
		name        string
		hand        []deck.Card // The player's hand
		run         []deck.Card // The run to play
		expectValid bool        // Whether the run is expected to be valid
	}{
		{
			name: "Valid run of deck.Four",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Spades, Points: 5},
				{Rank: deck.Four, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Five, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Six, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Joker, Suit: deck.AnySuit, Points: 50},
				{Rank: deck.Seven, Suit: deck.Spades, Points: 5},
				{Rank: deck.Eight, Suit: deck.Spades, Points: 5},
				{Rank: deck.Nine, Suit: deck.Spades, Points: 5},
				{Rank: deck.Ten, Suit: deck.Spades, Points: 10},
			},
			run: []deck.Card{
				{Rank: deck.Four, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Five, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Six, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Joker, Suit: deck.AnySuit, Points: 50},
			},
			expectValid: true,
		},
		{
			name: "Valid run of deck.Five",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Four, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Five, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Six, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Joker, Suit: deck.AnySuit, Points: 50},
				{Rank: deck.Seven, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Eight, Suit: deck.Spades, Points: 5},
				{Rank: deck.Nine, Suit: deck.Spades, Points: 5},
				{Rank: deck.Ten, Suit: deck.Spades, Points: 10},
				{Rank: deck.Jack, Suit: deck.Spades, Points: 10},
			},
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Four, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Five, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Six, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Seven, Suit: deck.Hearts, Points: 5},
			},
			expectValid: true,
		},
		{
			name: "Valid run of deck.Six",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Spades, Points: 5},
				{Rank: deck.Four, Suit: deck.Clubs, Points: 5},
				{Rank: deck.Five, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Six, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Joker, Suit: deck.AnySuit, Points: 50},
				{Rank: deck.Seven, Suit: deck.Spades, Points: 5},
				{Rank: deck.Eight, Suit: deck.Spades, Points: 5},
				{Rank: deck.Nine, Suit: deck.Spades, Points: 5},
				{Rank: deck.Ten, Suit: deck.Spades, Points: 10},
				{Rank: deck.Jack, Suit: deck.Spades, Points: 10},
				{Rank: deck.Queen, Suit: deck.Spades, Points: 10},
			},
			run: []deck.Card{
				{Rank: deck.Seven, Suit: deck.Spades, Points: 5},
				{Rank: deck.Eight, Suit: deck.Spades, Points: 5},
				{Rank: deck.Nine, Suit: deck.Spades, Points: 5},
				{Rank: deck.Ten, Suit: deck.Spades, Points: 10},
				{Rank: deck.Jack, Suit: deck.Spades, Points: 10},
				{Rank: deck.Queen, Suit: deck.Spades, Points: 10},
			},
			expectValid: true,
		},
		{
			name: "Valid run of deck.Seven",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Spades, Points: 5},
				{Rank: deck.Four, Suit: deck.Clubs, Points: 5},
				{Rank: deck.Five, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Six, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Joker, Suit: deck.AnySuit, Points: 50},
				{Rank: deck.Seven, Suit: deck.Spades, Points: 5},
				{Rank: deck.Eight, Suit: deck.Spades, Points: 5},
				{Rank: deck.Nine, Suit: deck.Spades, Points: 5},
				{Rank: deck.Ten, Suit: deck.Spades, Points: 10},
				{Rank: deck.Jack, Suit: deck.Spades, Points: 10},
				{Rank: deck.Queen, Suit: deck.Spades, Points: 10},
				{Rank: deck.King, Suit: deck.Spades, Points: 10},
			},
			run: []deck.Card{
				{Rank: deck.Seven, Suit: deck.Spades, Points: 5},
				{Rank: deck.Eight, Suit: deck.Spades, Points: 5},
				{Rank: deck.Nine, Suit: deck.Spades, Points: 5},
				{Rank: deck.Ten, Suit: deck.Spades, Points: 10},
				{Rank: deck.Jack, Suit: deck.Spades, Points: 10},
				{Rank: deck.Queen, Suit: deck.Spades, Points: 10},
				{Rank: deck.King, Suit: deck.Spades, Points: 10},
			},
			expectValid: true,
		},
		{
			name: "Invalid run with non-consecutive ranks",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Spades, Points: 5},
				{Rank: deck.Four, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Five, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Six, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Joker, Suit: deck.AnySuit, Points: 50},
				{Rank: deck.Seven, Suit: deck.Spades, Points: 5},
				{Rank: deck.Eight, Suit: deck.Spades, Points: 5},
				{Rank: deck.Nine, Suit: deck.Spades, Points: 5},
				{Rank: deck.Ten, Suit: deck.Spades, Points: 10},
			},
			run: []deck.Card{
				{Rank: deck.Three, Suit: deck.Spades, Points: 5},
				{Rank: deck.Five, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Six, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Joker, Suit: deck.AnySuit, Points: 50},
			},
			expectValid: false,
		},
		{
			name: "Invalid run not long enough",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Spades, Points: 5},
				{Rank: deck.Four, Suit: deck.Clubs, Points: 5},
				{Rank: deck.Five, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Six, Suit: deck.Hearts, Points: 5},
				{Rank: deck.Joker, Suit: deck.AnySuit, Points: 50},
				{Rank: deck.Seven, Suit: deck.Spades, Points: 5},
				{Rank: deck.Eight, Suit: deck.Spades, Points: 5},
				{Rank: deck.Nine, Suit: deck.Spades, Points: 5},
				{Rank: deck.Ten, Suit: deck.Spades, Points: 10},
			},
			run: []deck.Card{
				{Rank: deck.Seven, Suit: deck.Spades, Points: 5},
				{Rank: deck.Eight, Suit: deck.Spades, Points: 5},
				{Rank: deck.Nine, Suit: deck.Spades, Points: 5},
				{Rank: deck.Ten, Suit: deck.Spades, Points: 10},
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
			game = gameplay.NewGame(playerNames)
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

				if len(game.DownedSets[game.Players[0].Name]) != 1 {
					t.Errorf("Expected 1 downed set for player, got %d", len(game.DownedSets[game.Players[0].Name]))
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
	game := gameplay.NewGame(playerNames)

	// Define test cases
	testCases := []struct {
		name        string
		hand        []deck.Card // The player's hand
		book        []deck.Card // The book to play
		expectValid bool        // Whether the book is expected to be valid
	}{
		{
			name: "Valid book with deck.Three cards of the same rank",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Three, Suit: deck.Spades},
				{Rank: deck.Three, Suit: deck.Diamonds},
				{Rank: deck.Four, Suit: deck.Clubs},
			},
			book: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Three, Suit: deck.Spades},
				{Rank: deck.Three, Suit: deck.Diamonds},
			},
			expectValid: true,
		},
		{
			name: "Invalid book with mixed ranks",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Four, Suit: deck.Spades},
				{Rank: deck.Five, Suit: deck.Diamonds},
				{Rank: deck.Six, Suit: deck.Clubs},
			},
			book: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Four, Suit: deck.Spades},
				{Rank: deck.Five, Suit: deck.Diamonds},
			},
			expectValid: false,
		},
		{
			name: "Invalid book with less than deck.Three cards",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Three, Suit: deck.Spades},
				{Rank: deck.Four, Suit: deck.Diamonds},
			},
			book: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Three, Suit: deck.Spades},
			},
			expectValid: false,
		},
		{
			name: "Valid book with one wildcard",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Three, Suit: deck.Spades},
				{Rank: deck.Joker, Suit: deck.AnySuit},
				{Rank: deck.Four, Suit: deck.Clubs},
			},
			book: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Three, Suit: deck.Spades},
				{Rank: deck.Joker, Suit: deck.AnySuit},
			},
			expectValid: true,
		},
		{
			name: "Invalid book with deck.Two wildcards",
			hand: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Joker, Suit: deck.AnySuit},
				{Rank: deck.Joker, Suit: deck.AnySuit},
				{Rank: deck.Four, Suit: deck.Clubs},
			},
			book: []deck.Card{
				{Rank: deck.Three, Suit: deck.Hearts},
				{Rank: deck.Joker, Suit: deck.AnySuit},
				{Rank: deck.Joker, Suit: deck.AnySuit},
			},
			expectValid: false,
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange: Set up the player's hand
			game = gameplay.NewGame(playerNames)
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

				if len(game.DownedSets[game.Players[0].Name]) != 1 {
					t.Errorf("Expected 1 downed set for player, got %d", len(game.DownedSets[game.Players[0].Name]))
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
