package lobby

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmiller-57/Push/push-backend/internal/http/middleware"
)

type createReq struct {
	Name string `json:"name"`
	MaxPlayers int `json:"maxPlayers"`
}

func ListRoomsHandler(store *MemoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"rooms": store.List()})
	}
}

func CreateRoomHandler(store *MemoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createReq
		_ = json.NewDecoder(r.Body).Decode(&req)
		host := middleware.UserID(r)
		room := store.Create(req.Name, req.MaxPlayers, host)
		_ = json.NewEncoder(w).Encode(room)
	}
}

func JoinRoomHandler(store *MemoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user := middleware.UserID(r)
		if err := store.Join(id, user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}