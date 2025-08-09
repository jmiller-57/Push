package lobby

import (
	"errors"
	"sync"
)

type Room struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	MaxPlayers int      `json:"maxPlayers"`
	PlayerIDs  []string `json:"playerIds"`
}

type MemoryStore struct {
	mu    sync.Mutex
	rooms map[string]*Room
}

func NewMemoryStore() *MemoryStore { return &MemoryStore{rooms: map[string]*Room{}} }

func (s *MemoryStore) List() []*Room {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]*Room, 0, len(s.rooms))
	for _, r := range s.rooms {
		out = append(out, r)
	}
	return out
}

func (s *MemoryStore) Create(name string, max int, hostID string) *Room {
	s.mu.Lock()
	defer s.mu.Unlock()
	if name == "" {
		name = "Room"
	}
	if max <= 0 {
		max = 6
	}
	id := "r_" + name // replace with uuid
	r := &Room{ID: id, Name: name, MaxPlayers: max, PlayerIDs: []string{hostID}}
	s.rooms[id] = r
	return r
}

func (s *MemoryStore) Join(id, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, ok := s.rooms[id]
	if !ok {
		return errors.New("room not found")
	}
	if len(r.PlayerIDs) >= r.MaxPlayers {
		return errors.New("room full")
	}
	for _, pid := range r.PlayerIDs {
		if pid == userID {
			return nil
		}
	}
	r.PlayerIDs = append(r.PlayerIDs, userID)
	return nil
}
