package main

import (
	"context"
	"log"

	"github.com/amaan287/chess-backend/config"
	"github.com/amaan287/chess-backend/internal/migrations"
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

	if err = migrations.Run(context.Background(), store.DB()); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("migrations completed successfully")
}
