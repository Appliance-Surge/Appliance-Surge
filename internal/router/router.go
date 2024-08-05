package router

import (
    "net/http"

    "github.com/go-chi/chi/v5"

    "github.com/Appliance-Surge/Appliance-Surge/internal/config"
    "github.com/Appliance-Surge/Appliance-Surge/internal/handlers"
    "github.com/Appliance-Surge/Appliance-Surge/internal/storage"
    "github.com/Appliance-Surge/Appliance-Surge/internal/assets"
)

func NewRouter(cfg *config.Config, db *storage.DB) *chi.Mux {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        assetsData := map[string]interface{}{
            "CSSPaths": assets.CSSPaths("resources/js/main.js"),
            "JSPath":  assets.AssetPath("resources/js/main.js"),
        }
        handlers.HomeHandler(db, assetsData).ServeHTTP(w, r)
    })

    fileServer := http.FileServer(http.Dir("./static"))
    r.Handle("/static/*", http.StripPrefix("/static", fileServer))

    return r
}
