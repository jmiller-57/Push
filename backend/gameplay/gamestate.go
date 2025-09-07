package gameplay

import (
	"github.com/jmiller-57/Push/backend/gameplay/deck"
)

type GameState struct {
	Players        []Player
	CurrentTurnIdx int
	Deck           deck.Deck
	FaceUpCard     *deck.Card
	Round          RoundObjective
	HasStarted     bool
	HasEnded       bool
	DownedSets     map[string][]Play // Tracks sets played by each player
}

func NewGame(players []Player) GameState {
	deck := deck.NewDeck()

	for i := range players {
		hand := deck[:11]
		deck = deck[11:]
		players[i].Hand = hand
	}

	faceUp := deck[0]
	deck = deck[1:]

	return GameState{
		Players:        players,
		CurrentTurnIdx: 0,
		Deck:           deck,
		FaceUpCard:     &faceUp,
		Round:          RoundObjectives[0],
		HasStarted:     true,
		DownedSets:     make(map[string][]Play),
	}
}

func (g *GameState) TakeFaceUpCard() {
	player := &g.Players[g.CurrentTurnIdx]
	player.Hand = append(player.Hand, *g.FaceUpCard)
	g.FaceUpCard = nil
}

func (g *GameState) PushFaceUpCard() {
	leftIdx := (g.CurrentTurnIdx + 1) % len(g.Players)
	// Next player will receive this next turn
	nextPlayer := &g.Players[leftIdx]
	nextPlayer.Hand = append(nextPlayer.Hand, *g.FaceUpCard)
	nextPlayer.Hand = append(nextPlayer.Hand, g.Deck.DrawCard())

	g.Players[g.CurrentTurnIdx].Hand = append(g.Players[g.CurrentTurnIdx].Hand, g.Deck.DrawCard())
	g.FaceUpCard = nil
}

func (g *GameState) Discard(card deck.Card) {
	g.RemoveCardsFromHand([]deck.Card{card})
	g.FaceUpCard = &card
	g.AdvanceTurn()
}

func (g *GameState) AdvanceTurn() {
	g.CurrentTurnIdx = (g.CurrentTurnIdx + 1) % len(g.Players)
}

func (g *GameState) PlayBook(book []deck.Card) {
	player := &g.Players[g.CurrentTurnIdx]
	if ValidateBook(book) {
		g.RemoveCardsFromHand(book)
		g.DownedSets[player.Name] = append(g.DownedSets[player.Name],
			Play{
				Cards: book,
				Set: SetRequirement{
					Type:      Book,
					MinLength: 3}})
	} else {
		panic("Invalid book played")
	}
	g.AdvanceTurn()
}

func (g *GameState) PlayRun(run []deck.Card, suit deck.Suit, length int8) {
	if g.Round.RoundNumber == 1 ||
		g.Round.RoundNumber == 3 {
		panic("Runs cannot be played in rounds 1 or 3")
	}

	runLenReq := g.Round.Sets[RUN_INDEX].MinLength
	// Validate the run against the current round's requirements
	if len(run) < int(runLenReq) {
		panic("Run is too short")
	}

	player := &g.Players[g.CurrentTurnIdx]
	if ValidateRun(run, suit, length) {
		g.RemoveCardsFromHand(run)
		g.DownedSets[player.Name] = append(g.DownedSets[player.Name],
			Play{
				Cards: run,
				Set: SetRequirement{
					Type:      Run,
					MinLength: runLenReq}})
	} else {
		panic("Invalid run played")
	}
	g.AdvanceTurn()
}

func (g *GameState) RemoveCardsFromHand(cards []deck.Card) {
	player := &g.Players[g.CurrentTurnIdx]
	for _, card := range cards {
		for i, c := range player.Hand {
			if c.Equals(card) {
				player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
				break
			}
		}
	}
}
