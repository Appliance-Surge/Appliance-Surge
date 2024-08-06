package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Appliance-Surge/Appliance-Surge/internal/assets"
	"github.com/Appliance-Surge/Appliance-Surge/internal/config"
	"github.com/Appliance-Surge/Appliance-Surge/internal/router"
	"github.com/Appliance-Surge/Appliance-Surge/internal/server"
	"github.com/Appliance-Surge/Appliance-Surge/internal/storage"
)

// Welcome To Appliance Surge!
// This file Spins up the server, database, and routing!
//
// Since: 0.0.1
func main() {

	// Load the assets manifest created by Vite
	//
	// Since: 0.0.1
	assets.LoadManifest()

	// Load the config
	//
	// Since: 0.0.1
	cfg := config.LoadConfig()

	// Create the database connection
	//
	// Since: 0.0.1
	db := startDBConnection(cfg)

	// Create the router
	//
	// Since 0.0.1
	r := router.NewRouter(cfg, db)

	// Spin up the server
	//
	// Since 0.0.1
	startServer(cfg, r)
}

// Create database connection
//
// Params:
// - cfg (*config.Config): Configurations for the server
//
// Returns *storage.DB: Database connection
//
// Since 0.0.1
func startDBConnection(cfg *config.Config) *storage.DB {
	db, err := storage.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	return db
}

// Create a server instance
//
// Params:
// - cfg (*config.Config): Configurations for the server
// - r (*chi.Mux): The web server router
//
// Returns: *http.Server
//
// Since: 0.0.1
func startServer(cfg *config.Config, r *chi.Mux) *http.Server {
	srv := server.NewServer(cfg, r)
	log.Println("Starting server...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	return srv
}
