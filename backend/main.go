package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmiller-57/Push/backend/handlers"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	db := InitDB("push.db")
	defer db.Close()

	userHandler := handlers.NewUserHandler(db)
	roomHandler := handlers.NewRoomHandler(db)
	gameHandler := handlers.NewGameHandler(db)

	r := mux.NewRouter()
	r.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/login", userHandler.Login).Methods("POST")
	r.Handle("/api/lobby", handlers.JWTAuthMiddleware(http.HandlerFunc(roomHandler.CreateRoom))).Methods("POST")
	r.Handle("/api/lobby/rooms/join", handlers.JWTAuthMiddleware(http.HandlerFunc(roomHandler.JoinRoom))).Methods("POST")
	r.Handle("/api/lobby/rooms/list", handlers.JWTAuthMiddleware(http.HandlerFunc(roomHandler.ListRooms))).Methods("GET")
	r.Handle("/api/lobby/rooms/{id}", handlers.JWTAuthMiddleware(http.HandlerFunc(roomHandler.RoomDetails))).Methods("GET")
	r.Handle("/api/lobby/rooms/{id}/start", handlers.JWTAuthMiddleware(http.HandlerFunc(gameHandler.StartGame))).Methods("POST")
	r.Handle("/api/lobby/rooms/{id}/game", handlers.JWTAuthMiddleware(http.HandlerFunc(gameHandler.GetGameState))).Methods("GET")
	r.Handle("/api/lobby/rooms/{id}/play", handlers.JWTAuthMiddleware(http.HandlerFunc(gameHandler.PlayCards))).Methods("POST")

	log.Println("Server started listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(r)))
}
