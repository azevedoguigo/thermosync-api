package main

import (
	"log"
	"net/http"

	"github.com/azevedoguigo/thermosync-api/config"
	"github.com/azevedoguigo/thermosync-api/internal/handler"
	"github.com/azevedoguigo/thermosync-api/internal/repository"
	"github.com/azevedoguigo/thermosync-api/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
	})

	log.Println("Server is running in port: 3000")
	http.ListenAndServe(":3000", router)
}
