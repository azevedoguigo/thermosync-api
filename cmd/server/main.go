package main

import (
	"log"
	"net/http"

	"github.com/azevedoguigo/thermosync-api/config"
	"github.com/azevedoguigo/thermosync-api/internal/handler"
	authMiddleware "github.com/azevedoguigo/thermosync-api/internal/middleware"
	"github.com/azevedoguigo/thermosync-api/internal/repository"
	"github.com/azevedoguigo/thermosync-api/internal/service"
	"github.com/azevedoguigo/thermosync-api/internal/websocket"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.With(authMiddleware.AuthMiddleware).Get("/{id}", userHandler.FindUserByID)
	})

	router.Route("/auth", func(r chi.Router) {
		r.Post("/", authHandler.Login)
	})

	router.Get("/ws", handler.Websocket)

	go websocket.HandleMessages()

	log.Println("Server is running in port: 3000")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatalf("Error to start server: %s", err)
	}
}
