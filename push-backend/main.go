package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmiller-57/Push/push-backend/handlers"
)

func main() {
	db := InitDB("push.db")
	defer db.Close()

	userHandler := handlers.NewUserHandler(db)
	roomHandler := handlers.NewRoomHandler(db)

	r := mux.NewRouter()
	r.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/login", userHandler.Login).Methods("POST")
	r.Handle("/api/rooms", handlers.JWTAuthMiddleware(http.HandlerFunc(roomHandler.CreateRoom))).Methods("POST")
	r.Handle("/api/rooms/join", handlers.JWTAuthMiddleware(http.HandlerFunc(roomHandler.JoinRoom))).Methods("POST")
	r.HandleFunc("/api/rooms/list", roomHandler.ListRooms).Methods("GET")
	r.HandleFunc("/api/rooms/{id}", roomHandler.RoomDetails).Methods("GET")

	log.Println("Server started listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
