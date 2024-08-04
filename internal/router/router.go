package router

import (
    "database/sql"

    "github.com/go-chi/chi/v5"

    "github.com/Appliance-Surge/Appliance-Surge/internal/config"
    "github.com/Appliance-Surge/Appliance-Surge/internal/handlers"
)

func NewRouter(cfg *config.Config, db *sql.DB) *chi.Mux {
    r := chi.NewRouter()

    r.Get("/posts", handlers.GetPosts(db))
    r.Post("/posts", handlers.CreatePost(db))

    return r
}
