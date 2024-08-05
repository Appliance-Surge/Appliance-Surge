package main

import (
    "log"

    "github.com/Appliance-Surge/Appliance-Surge/internal/config"
    "github.com/Appliance-Surge/Appliance-Surge/internal/router"
    "github.com/Appliance-Surge/Appliance-Surge/internal/server"
    "github.com/Appliance-Surge/Appliance-Surge/internal/storage"
    "github.com/Appliance-Surge/Appliance-Surge/internal/assets"
)

func main() {

    assets.LoadManifest()

    cfg := config.LoadConfig()
    db, err := storage.NewDB(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    r := router.NewRouter(cfg, db)
    srv := server.NewServer(cfg, r)

    log.Println("Starting server...")
    if err := srv.ListenAndServe(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
