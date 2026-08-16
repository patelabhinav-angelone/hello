package main

import (
	"log"

	"hello/config"
	"hello/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on existing environment variables")
	}

	config.ConnectDatabase()

	r := router.SetupRouter()
	log.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
