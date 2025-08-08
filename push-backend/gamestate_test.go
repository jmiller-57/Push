package main

import (
    "testing"
)
var g GameState
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
		hand        []Card // The player's hand
		run         []Card // The run to play
		expectValid bool   // Whether the run is expected to be valid
	}{
		  {
				name: "Valid run of four",
				hand: []Card{
					{Rank: three, Suit: spades, Points: 5, PossibleValues: []int8{3}},
					{Rank: four, Suit: hearts, Points: 5, PossibleValues: []int8{4}},
					{Rank: five, Suit: hearts, Points: 5, PossibleValues: []int8{5}},
					{Rank: six, Suit: hearts, Points: 5, PossibleValues: []int8{6}},
					{Rank: joker, Suit: anySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
					{Rank: seven, Suit: spades, Points: 5, PossibleValues: []int8{7}},
					{Rank: eight, Suit: spades, Points: 5, PossibleValues: []int8{8}},
					{Rank: nine, Suit: spades, Points: 5, PossibleValues: []int8{9}},
					{Rank: ten, Suit: spades, Points: 10, PossibleValues: []int8{10}},
				},
				run: []Card{
					{Rank: four, Suit: hearts, Points: 5, PossibleValues: []int8{4}},
					{Rank: five, Suit: hearts, Points: 5, PossibleValues: []int8{5}},
					{Rank: six, Suit: hearts, Points: 5, PossibleValues: []int8{6}},
					{Rank: joker, Suit: anySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
					},
				expectValid: true,
			},
			{
				name: "Valid run of five",
				hand: []Card{
					{Rank: three, Suit: hearts, Points: 5, PossibleValues: []int8{3}},
					{Rank: four, Suit: hearts, Points: 5, PossibleValues: []int8{4}},
					{Rank: five, Suit: hearts, Points: 5, PossibleValues: []int8{5}},
					{Rank: six, Suit: hearts, Points: 5, PossibleValues: []int8{6}},
					{Rank: joker, Suit: anySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
					{Rank: seven, Suit: hearts, Points: 5, PossibleValues: []int8{7}},
					{Rank: eight, Suit: spades, Points: 5, PossibleValues: []int8{8}},
					{Rank: nine, Suit: spades, Points: 5, PossibleValues: []int8{9}},
					{Rank: ten, Suit: spades, Points: 10, PossibleValues: []int8{10}},
					{Rank: jack, Suit: spades, Points: 10, PossibleValues: []int8{11}},
				},
				run: []Card{
					{Rank: three, Suit: hearts, Points: 5, PossibleValues: []int8{3}},
					{Rank: four, Suit: hearts, Points: 5, PossibleValues: []int8{4}},
					{Rank: five, Suit: hearts, Points: 5, PossibleValues: []int8{5}},
					{Rank: six, Suit: hearts, Points: 5, PossibleValues: []int8{6}},
					{Rank: seven, Suit: hearts, Points: 5, PossibleValues: []int8{7}},
				},
				expectValid: true,
			},
			{
				name: "Valid run of six",
				hand: []Card{
					{Rank: three, Suit: spades, Points: 5, PossibleValues: []int8{3}},
					{Rank: four, Suit: clubs, Points: 5, PossibleValues: []int8{4}},
					{Rank: five, Suit: hearts, Points: 5, PossibleValues: []int8{5}},
					{Rank: six, Suit: hearts, Points: 5, PossibleValues: []int8{6}},
					{Rank: joker, Suit: anySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
					{Rank: seven, Suit: spades, Points: 5, PossibleValues: []int8{7}},
					{Rank: eight, Suit: spades, Points: 5, PossibleValues: []int8{8}},
					{Rank: nine, Suit: spades, Points: 5, PossibleValues: []int8{9}},
					{Rank: ten, Suit: spades, Points: 10, PossibleValues: []int8{10}},
					{Rank: jack, Suit: spades, Points: 10, PossibleValues: []int8{11}},
					{Rank: queen, Suit: spades, Points: 10, PossibleValues: []int8{12}},
				},
				run: []Card{
					{Rank: seven, Suit: spades, Points: 5, PossibleValues: []int8{7}},
					{Rank: eight, Suit: spades, Points: 5, PossibleValues: []int8{8}},
					{Rank: nine, Suit: spades, Points: 5, PossibleValues: []int8{9}},
					{Rank: ten, Suit: spades, Points: 10, PossibleValues: []int8{10}},
					{Rank: jack, Suit: spades, Points: 10, PossibleValues: []int8{11}},
					{Rank: queen, Suit: spades, Points: 10, PossibleValues: []int8{12}},
				},
				expectValid: true,
			},
			{
				name: "Valid run of seven",
				hand: []Card{
					{Rank: three, Suit: spades, Points: 5, PossibleValues: []int8{3}},
					{Rank: four, Suit: clubs, Points: 5, PossibleValues: []int8{4}},
					{Rank: five, Suit: hearts, Points: 5, PossibleValues: []int8{5}},
					{Rank: six, Suit: hearts, Points: 5, PossibleValues: []int8{6}},
					{Rank: joker, Suit: anySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
					{Rank: seven, Suit: spades, Points: 5, PossibleValues: []int8{7}},
					{Rank: eight, Suit: spades, Points: 5, PossibleValues: []int8{8}},
					{Rank: nine, Suit: spades, Points: 5, PossibleValues: []int8{9}},
					{Rank: ten, Suit: spades, Points: 10, PossibleValues: []int8{10}},
					{Rank: jack, Suit: spades, Points: 10, PossibleValues: []int8{11}},
					{Rank: queen, Suit: spades, Points: 10, PossibleValues: []int8{12}},
					{Rank: king, Suit: spades, Points: 10, PossibleValues: []int8{13}},
				},
				run: []Card{
					{Rank: seven, Suit: spades, Points: 5, PossibleValues: []int8{7}},
					{Rank: eight, Suit: spades, Points: 5, PossibleValues: []int8{8}},
					{Rank: nine, Suit: spades, Points: 5, PossibleValues: []int8{9}},
					{Rank: ten, Suit: spades, Points: 10, PossibleValues: []int8{10}},
					{Rank: jack, Suit: spades, Points: 10, PossibleValues: []int8{11}},
					{Rank: queen, Suit: spades, Points: 10, PossibleValues: []int8{12}},
					{Rank: king, Suit: spades, Points: 10, PossibleValues: []int8{13}},
				},
				expectValid: true,
			},
			{
				name: "Invalid run with non-consecutive ranks",
				hand: []Card{
					{Rank: three, Suit: spades, Points: 5, PossibleValues: []int8{3}},
					{Rank: four, Suit: hearts, Points: 5, PossibleValues: []int8{4}},
					{Rank: five, Suit: hearts, Points: 5, PossibleValues: []int8{5}},
					{Rank: six, Suit: hearts, Points: 5, PossibleValues: []int8{6}},
					{Rank: joker, Suit: anySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
					{Rank: seven, Suit: spades, Points: 5, PossibleValues: []int8{7}},
					{Rank: eight, Suit: spades, Points: 5, PossibleValues: []int8{8}},
					{Rank: nine, Suit: spades, Points: 5, PossibleValues: []int8{9}},
					{Rank: ten, Suit: spades, Points: 10, PossibleValues: []int8{10}},
				},
				run: []Card{
					{Rank: three, Suit: spades, Points: 5, PossibleValues: []int8{3}},
					{Rank: five, Suit: hearts, Points: 5, PossibleValues: []int8{5}},
					{Rank: six, Suit: hearts, Points: 5, PossibleValues: []int8{6}},
					{Rank: joker, Suit: anySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
				},
				expectValid: false,
			},
			{
				name: "Invalid run not long enough",
				hand: []Card{
					{Rank: three, Suit: spades, Points: 5, PossibleValues: []int8{3}},
					{Rank: four, Suit: clubs, Points: 5, PossibleValues: []int8{4}},
					{Rank: five, Suit: hearts, Points: 5, PossibleValues: []int8{5}},
					{Rank: six, Suit: hearts, Points: 5, PossibleValues: []int8{6}},
					{Rank: joker, Suit: anySuit, Points: 50, PossibleValues: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
					{Rank: seven, Suit: spades, Points: 5, PossibleValues: []int8{7}},
					{Rank: eight, Suit: spades, Points: 5, PossibleValues: []int8{8}},
					{Rank: nine, Suit: spades, Points: 5, PossibleValues: []int8{9}},
					{Rank: ten, Suit: spades, Points: 10, PossibleValues: []int8{10}},
				},
				run: []Card{
					{Rank: seven, Suit: spades, Points: 5, PossibleValues: []int8{7}},
					{Rank: eight, Suit: spades, Points: 5, PossibleValues: []int8{8}},
					{Rank: nine, Suit: spades, Points: 5, PossibleValues: []int8{9}},
					{Rank: ten, Suit: spades, Points: 10, PossibleValues: []int8{10}},
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
			roundObj := RoundObjective{
				RoundNumber: roundIdx + 1,
				Sets: RoundObjectives[roundIdx].Sets,
			}
			game.Round = roundObj
			runLength := RoundObjectives[roundIdx].Sets[RUN_INDEX].MinLength

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
			hand        []Card // The player's hand
			book        []Card // The book to play
			expectValid bool   // Whether the book is expected to be valid
	}{
			{
					name: "Valid book with three cards of the same rank",
					hand: []Card{
							{Rank: three, Suit: hearts},
							{Rank: three, Suit: spades},
							{Rank: three, Suit: diamonds},
							{Rank: four, Suit: clubs},
					},
					book: []Card{
							{Rank: three, Suit: hearts},
							{Rank: three, Suit: spades},
							{Rank: three, Suit: diamonds},
					},
					expectValid: true,
			},
			{
					name: "Invalid book with mixed ranks",
					hand: []Card{
							{Rank: three, Suit: hearts},
							{Rank: four, Suit: spades},
							{Rank: five, Suit: diamonds},
							{Rank: six, Suit: clubs},
					},
					book: []Card{
							{Rank: three, Suit: hearts},
							{Rank: four, Suit: spades},
							{Rank: five, Suit: diamonds},
					},
					expectValid: false,
			},
			{
					name: "Invalid book with less than three cards",
					hand: []Card{
							{Rank: three, Suit: hearts},
							{Rank: three, Suit: spades},
							{Rank: four, Suit: diamonds},
					},
					book: []Card{
							{Rank: three, Suit: hearts},
							{Rank: three, Suit: spades},
					},
					expectValid: false,
			},
			{
					name: "Valid book with one wildcard",
					hand: []Card{
							{Rank: three, Suit: hearts},
							{Rank: three, Suit: spades},
							{Rank: joker, Suit: anySuit},
							{Rank: four, Suit: clubs},
					},
					book: []Card{
							{Rank: three, Suit: hearts},
							{Rank: three, Suit: spades},
							{Rank: joker, Suit: anySuit},
					},
					expectValid: true,
			},
			{
					name: "Invalid book with two wildcards",
					hand: []Card{
							{Rank: three, Suit: hearts},
							{Rank: joker, Suit: anySuit},
							{Rank: joker, Suit: anySuit},
							{Rank: four, Suit: clubs},
					},
					book: []Card{
							{Rank: three, Suit: hearts},
							{Rank: joker, Suit: anySuit},
							{Rank: joker, Suit: anySuit},
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