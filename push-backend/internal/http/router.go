package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jmiller-57/Push/push-backend/internal/auth"
	"github.com/jmiller-57/Push/push-backend/internal/http/middleware"
	"github.com/jmiller-57/Push/push-backend/internal/lobby"
	"github.com/jmiller-57/Push/push-backend/internal/ws"
)

func New(jwtSecret, frontendOrigin string ) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:			[]string{frontendOrigin},
		AllowedMethods:			[]string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:			[]string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials:	true,
	}))

	authSvc := auth.NewService(jwtSecret)
	roomStore := lobby.NewMemoryStore()
	hub := ws.NewHub(roomStore, authSvc)
	go hub.Run()

	r.Route("/api", func(api chi.Router) {
		api.Post("/register", auth.RegisterHandler(authSvc))
		api.Post("/login", auth.LoginHandler(authSvc))

		api.Group(func(pr chi.Router) {
			pr.Use(middleware.JWT(authSvc))
			pr.Get("/rooms", lobby.ListRoomsHandler(roomStore))
			pr.Post("/rooms", lobby.CreateRoomHandler(roomStore))
			pr.Post("/rooms/{id}/join", lobby.JoinRoomHandler(roomStore))
		})
	})

	r.Get("/ws/games/{id}", ws.GameWSHandler(hub, authSvc))

	return r
}