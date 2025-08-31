package gameplay

import "github.com/google/uuid"

type PlayerID uuid.UUID

type Player struct {
	ID     PlayerID
	Name   string
	Hand   []Card
	Score  int
	IsDown bool
}
