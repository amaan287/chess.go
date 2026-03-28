package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amaan287/chess-backend/config"
	"github.com/amaan287/chess-backend/constants"
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
		log.Print("Failed to get env")
	}
	port := fmt.Sprintf(":%v", env.ServerPort)
	err = http.ListenAndServe(string(port), nil)

}
