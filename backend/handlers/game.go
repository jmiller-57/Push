package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/jmill-57/Push/backend/gameplay"
)

type GameHandler struct {
	DB *sql.DB
}

func NewGameHandler(db *sql.DB) *GameHandler {
	return &GameHandler{DB: db}
}

func (h *GameHandler) StartGame(w http.ResponseWriter, r *http.Request) {
	
}
