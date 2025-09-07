package gameplay

import (
	"github.com/jmiller-57/Push/backend/gameplay/deck"
)

type Player struct {
	ID     int64
	Name   string
	Hand   []deck.Card
	Score  int
	IsDown bool
}
