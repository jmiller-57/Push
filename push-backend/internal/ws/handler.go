package ws

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	UserID string
	Conn *websocket.Conn
	send chan Envelope
	room *RoomHub
}

func GameWSHandler(h *Hub, auth jwtVerifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomID := chi.URLParam(r, "id")
		token := r.URL.Query().Get("token")
		userID, err := auth.Verify(token)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		room := h.room(roomID)
		c := &Client{
			UserID: userID,
			Conn: conn,
			send: make(chan Envelope, 16),
			room: room,
		}
		room.Clients[userID] = c

		// send initial state (placeholder)
		state := fakeState(h, roomID, userID)
		c.send <- Envelope{Type: "STATE", Payload: mustJSON(state)}

		go c.writeLoop()
		c.readLoop()
	}
}

func (c *Client) readLoop() {
	defer func() {
		delete(c.room.Clients, c.UserID)
		c.Conn.Close()
	}()

	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		log.Printf("recv from %s: %s", c.UserID, string(data))
		c.room.Broadcast <- Envelope{Type: "STATE", Payload: mustJSON(map[string]any{
			"roomId": c.room.ID,
		})}
	}
}

func (c *Client) writeLoop() {
	for msg := range c.send {
		_ = c.Conn.WriteJSON(msg)
	}
}

func mustJSON(v any) []byte {
	b, _ := json.Marshal(v)
	return b
}

func fakeState(h *Hub, roomID, me string) map[string]any {
	var players []map[string]any
	for _, r := range h.Store().List() {
		if r.ID != roomID {
			continue
		}
		for _, pid := range r.PlayerIDs {
			players = append(players, map[string]any{
				"id": pid,
				"name": pid,
				"isMe": pid == me,
				"handCount": 11,
			})
		}
	}
	current := ""
	if len(players) > 0 {
		if id, ok := players[0]["id"].(string); ok {
			current = id
		}
	}
	return map[string]any{
		"roomId": roomID,
		"players": players,
		"currentPlayerId": current,
		"faceUpCard": nil,
		"drawPileCount": 108,
	}
}