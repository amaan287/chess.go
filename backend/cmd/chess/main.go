package main

import (
	"fmt"
	"github.com/amaan287/chess-backend/config"
	"github.com/amaan287/chess-backend/constants"
	"github.com/amaan287/chess-backend/internal/auth"
	"github.com/amaan287/chess-backend/internal/handler"
	"github.com/amaan287/chess-backend/internal/repository"
	"github.com/amaan287/chess-backend/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	store, err := config.NewPostgresStore()
	if err != nil {
		log.Fatalf("failed to initialize postgres store: %v", err)
	}
	defer func() {
		if closeErr := store.Close(); closeErr != nil {
			log.Printf("failed to close postgres store: %v", closeErr)
		}
	}()
	log.Println("postgres store initialized")

	env, err := constants.GetEnv()
	if err != nil {
		log.Fatalf("failed to get env: %v", err)
	}

	accessTTL := time.Duration(env.AccessTokenMinutes) * time.Minute
	refreshTTL := time.Duration(env.RefreshTokenHours) * time.Hour
	tokenManager := auth.NewManager(env.JWTSecret, accessTTL, refreshTTL)

	userRepo := repository.NewPostgresUserRepository(store.DB())
	userService := service.NewUserService(userRepo, tokenManager)
	userHandler := handler.NewUserHandler(userService)

	router := mux.NewRouter()
	userHandler.RegisterRoutes(router)

	port := fmt.Sprintf(":%d", env.ServerPort)
	log.Printf("server listening on %s", port)

	if err = http.ListenAndServe(port, router); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
