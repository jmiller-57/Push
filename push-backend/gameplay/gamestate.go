package gameplay

import "github.com/google/uuid"

type GameState struct {
	Players        []Player
	CurrentTurnIdx int
	Deck           Deck
	FaceUpCard     *Card
	Round          RoundObjective
	HasStarted     bool
	HasEnded       bool
	DownedSets     map[PlayerID][]Play // Tracks sets played by each player
}

func (g *GameState) NewGame(playerNames []string) GameState {
	deck := NewDeck()

	players := make([]Player, len(playerNames))
	for i, name := range playerNames {
		hand := deck[:11]
		deck = deck[11:]
		players[i] = Player{ID: PlayerID(uuid.New()), Name: name, Hand: hand}
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
		DownedSets:     make(map[PlayerID][]Play),
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

func (g *GameState) Discard(card Card) {
	g.RemoveCardsFromHand([]Card{card})
	g.FaceUpCard = &card
	g.AdvanceTurn()
}

func (g *GameState) AdvanceTurn() {
	g.CurrentTurnIdx = (g.CurrentTurnIdx + 1) % len(g.Players)
}

func (g *GameState) PlayBook(book []Card) {
	player := &g.Players[g.CurrentTurnIdx]
	if ValidateBook(book) {
		g.RemoveCardsFromHand(book)
		g.DownedSets[player.ID] = append(g.DownedSets[player.ID],
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

func (g *GameState) PlayRun(run []Card, suit Suit, length int8) {
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
		g.DownedSets[player.ID] = append(g.DownedSets[player.ID],
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

func (g *GameState) RemoveCardsFromHand(cards []Card) {
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
