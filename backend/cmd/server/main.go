package main

import (
	"log"

	"github.com/the-sandwich/backend/internal/api"
	"github.com/the-sandwich/backend/internal/db"
	"github.com/the-sandwich/backend/internal/models"
	"github.com/the-sandwich/backend/internal/redis"
)

func main() {
	// Setup Postgres
	// In production, load from ENV.
	dsn := "host=localhost user=postgres password=postgres dbname=sandwich port=5432 sslmode=disable"
	if err := db.ConnectDB(dsn); err != nil {
		log.Printf("Warning: Failed to connect to DB: %v", err)
	} else {
		models.MigrateUsers(db.DB)
	}

	if err := redis.ConnectRedis("localhost:6379", "", 0); err != nil {
		log.Printf("Warning: Failed to connect to Redis: %v", err)
	}

	r := api.SetupRouter()
	log.Println("Starting server on :8080")
	r.Run(":8080")
}
