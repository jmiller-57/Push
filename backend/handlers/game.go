package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmiller-57/Push/backend/gameplay"
)

type GameHandler struct {
	DB *sql.DB
}

func NewGameHandler(db *sql.DB) *GameHandler {
	return &GameHandler{DB: db}
}

func (h *GameHandler) StartGame(w http.ResponseWriter, r *http.Request) {
	// Get room ID from URL
	vars := mux.Vars(r)
	roomIDStr := vars["id"]
	roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid room id", http.StatusBadRequest)
		return
	}

	var stateBytes []byte
	err = h.DB.QueryRow("SELECT state FROM games WHERE room_id = ?", roomID).Scan(&stateBytes)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(stateBytes)
		return
	} else if err != sql.ErrNoRows {
		http.Error(w, "error checking for existing game", http.StatusInternalServerError)
		return
	}

	// Get room members
	rows, err := h.DB.Query(`
		SELECT u.username
		FROM room_members rm
		JOIN users u
		ON u.id = rm.user_id
		WHERE room_id = ?`, roomID)
	if err != nil {
		http.Error(w, "could not get room members", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var players []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			http.Error(w, "error reading members", http.StatusInternalServerError)
			return
		}
		players = append(players, username)
	}
	if len(players) < 2 {
		http.Error(w, "need at least 2 players to start", http.StatusBadRequest)
		return
	}

	gameState := gameplay.NewGame(players)

	stateBytes, err = json.Marshal(gameState)
	if err != nil {
		http.Error(w, "could not serialize game state", http.StatusInternalServerError)
		return
	}

	_, err = h.DB.Exec("INSERT INTO games (room_id, state) VALUES (?, ?) ON CONFLICT(room_id) DO UPDATE SET state = ?", roomID, stateBytes, stateBytes)
	if err != nil {
		http.Error(w, "could not save game state", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(stateBytes)
}

func (h *GameHandler) GetGameState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomIDStr := vars["id"]
	roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid room id", http.StatusBadRequest)
		return
	}

	var stateBytes []byte
	err = h.DB.QueryRow("SELECT state FROM games WHERE room_id = ?", roomID).Scan(&stateBytes)
	if err != nil {
		if err == sql.ErrNoRows {
			// Game not started yet
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"gameStarted": false}`))
			return
		}
		http.Error(w, "could not get game state", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(stateBytes)
}
