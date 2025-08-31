package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type RoomHandler struct {
	DB *sql.DB
}

func NewRoomHandler(db *sql.DB) *RoomHandler {
	return &RoomHandler{DB: db}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
    var req struct {
        RoomName  string `json:"roomname"`
        CreatorID int    `json:"creator_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.RoomName == "" {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    res, err := h.DB.Exec("INSERT INTO rooms (roomname, creator_id) VALUES (?, ?)", req.RoomName, req.CreatorID)
    if err != nil {
        http.Error(w, "could not create room", http.StatusInternalServerError)
        return
    }
    id, _ := res.LastInsertId()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "id":         id,
        "roomname":   req.RoomName,
        "creator_id": req.CreatorID,
    })
}

func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
    var req struct {
        RoomID int64 `json:"room_id"`
        UserID int64 `json:"user_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    _, err := h.DB.Exec("INSERT INTO room_members (room_id, user_id) VALUES (?, ?)", req.RoomID, req.UserID)
    if err != nil {
        http.Error(w, "could not join room", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *RoomHandler) ListRooms(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, roomname, creator_id FROM rooms")
	if err != nil {
		http.Error(w, "could not list rooms", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Room struct {
		ID        int64  `json:"id"`
		RoomName  string `json:"roomname"`
		CreatorID int64  `json:"creator_id"`
	}

	var rooms []Room
	for rows.Next() {
		var room Room
		if err := rows.Scan(&room.ID, &room.RoomName, &room.CreatorID); err != nil {
			http.Error(w, "error reading rooms", http.StatusInternalServerError)
			return
		}
		rooms = append(rooms, room)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

func (h *RoomHandler) RoomDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomIDStr := vars["id"]
	roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil || roomIDStr == "" {
		http.Error(w, "invalid or missing room id", http.StatusBadRequest)
		return
	}

	// Get room info
	var room struct {
		ID        int64  `json:"id"`
		RoomName  string `json:"roomname"`
		CreatorID int64  `json:"creator_id"`
	}
	err = h.DB.QueryRow("SELECT id, roomname, creator_id FROM rooms WHERE id = ?", roomID).Scan(&room.ID, &room.RoomName, &room.CreatorID)
	if err != nil {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	var creatorStr string
	err = h.DB.QueryRow("SELECT username FROM users WHERE id = ?", room.CreatorID).Scan(&creatorStr)
	if err != nil {
		http.Error(w, "error retrieving creator username", http.StatusInternalServerError)
		return
	}

	rows, err := h.DB.Query(`
		SELECT u.id, u.username
		FROM room_members rm
		JOIN users u ON u.id = rm.user_id
		WHERE rm.room_id = ?`, roomID)
	if err != nil {
		http.Error(w, "could not get members", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Member struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
	}
	var members []Member
	for rows.Next() {
		var m Member
		if err := rows.Scan(&m.ID, &m.Username); err != nil {
			http.Error(w, "error reading members", http.StatusInternalServerError)
			return
		}
		members = append(members, m)
	}

	// Respond with room info and members
	resp := map[string]interface{}{
		"id":      room.ID,
		"name":    room.RoomName,
		"creator": creatorStr,
		"members": members,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
