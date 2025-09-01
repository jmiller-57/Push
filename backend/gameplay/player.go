package gameplay

import (
	"github.com/jmiller-57/Push/backend/gameplay/deck"
)

type PlayerID int64

type Player struct {
	Name   string
	Hand   []deck.Card
	Score  int
	IsDown bool
}
