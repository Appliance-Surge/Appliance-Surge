package main

import (
	"log"

	"github.com/Appliance-Surge/Appliance-Surge/db/factories"
	"github.com/Appliance-Surge/Appliance-Surge/internal/storage"
)

func main() {
	// Load configuration and initialize the database connection
	db, err := storage.NewDB("postgres://postgres:postgres@appliance_surge_db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create 10 users
	for i := 0; i < 10; i++ {
		if err := factories.CreateUserFactory(db); err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
	}

	log.Println("10 Users created successfully")

	// Create 10 posts
	for i := 0; i < 10; i++ {
		if err := factories.CreatePostFactory(db); err != nil {
			log.Fatalf("Failed to create post: %v", err)
		}
	}

	log.Println("10 Posts created successfully")
}
