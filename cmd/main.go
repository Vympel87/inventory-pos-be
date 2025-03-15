package main

import (
	"log"

	"github.com/Vympel87/inventory-pos-be/config"
	"github.com/Vympel87/inventory-pos-be/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Inisialisasi logger, database, redis
	config.LoadEnv()
	config.InitLogger()
	config.ConnectDatabase()
	config.ConnectRedis()

	// Start server
	app.Start()
}
