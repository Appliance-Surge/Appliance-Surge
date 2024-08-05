package router

import (
    "net/http"

    "github.com/go-chi/chi/v5"

    "github.com/Appliance-Surge/Appliance-Surge/internal/config"
    "github.com/Appliance-Surge/Appliance-Surge/internal/handlers"
    "github.com/Appliance-Surge/Appliance-Surge/internal/storage"
)

func NewRouter(cfg *config.Config, db *storage.DB) *chi.Mux {
    r := chi.NewRouter()

    r.Get("/", handlers.HomeHandler(db))

    fileServer := http.FileServer(http.Dir("./static"))
    r.Handle("/static/*", http.StripPrefix("/static", fileServer))

    fileServer = http.FileServer(http.Dir("./public"))
    r.Handle("/public/*", http.StripPrefix("/public", fileServer))

    return r
}
