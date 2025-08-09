package ws

import (
	"encoding/json"
	"sync"

	"github.com/jmiller-57/Push/push-backend/internal/lobby"
)

type jwtVerifier interface{ Verify(token string) (string, error) }

type Hub struct {
	mu sync.Mutex
	rooms map[string]*RoomHub
	store *lobby.MemoryStore
	auth jwtVerifier
}

func NewHub(store *lobby.MemoryStore, auth jwtVerifier) *Hub {
	return &Hub{rooms: map[string]*RoomHub{}, store: store, auth: auth}
}

func (h *Hub) Run() {}

func (h *Hub) Store() *lobby.MemoryStore { return h.store }

type Envelope struct {
	Type string `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type RoomHub struct {
	ID string
	Clients map[string]*Client
	Broadcast chan Envelope
}

func (h *Hub) room(id string) *RoomHub {
	h.mu.Lock()
	defer h.mu.Unlock()
	if r, ok := h.rooms[id]; ok {
		return r
	}
	r := &RoomHub{
		ID: id,
		Clients: map[string]*Client{},
		Broadcast: make(chan Envelope, 64),
	}
	h.rooms[id] = r
	go r.loop()
	return r
}

func (r *RoomHub) loop() {
	for msg := range r.Broadcast {
		for _, c := range r.Clients {
			c.send <- msg
		}
	}
}