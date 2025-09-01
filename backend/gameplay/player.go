package gameplay

import (
	"github.com/google/uuid"
	"github.com/jmiller-57/Push/backend/gameplay/deck"
)

type PlayerID uuid.UUID

type Player struct {
	ID     PlayerID
	Name   string
	Hand   []deck.Card
	Score  int
	IsDown bool
}
