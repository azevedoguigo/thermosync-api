package main

import (
	"log"
	"net/http"

	"github.com/azevedoguigo/thermosync-api/config"
	"github.com/azevedoguigo/thermosync-api/internal/handler"
	"github.com/azevedoguigo/thermosync-api/internal/repository"
	"github.com/azevedoguigo/thermosync-api/internal/service"
	"github.com/azevedoguigo/thermosync-api/internal/websocket"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	authHandler := handler.NewAuthHandler(userService)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.FindUserByID)
	})

	router.Route("/auth", func(r chi.Router) {
		r.Post("/", authHandler.Login)
	})

	router.Get("/ws", handler.Websocket)

	go websocket.HandleMessages()

	log.Println("Server is running in port: 3000")
	http.ListenAndServe(":3000", router)
}
