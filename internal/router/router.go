// This file handles all routing for the web server.
//
// Since: 0.0.1
package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Appliance-Surge/Appliance-Surge/internal/config"
	"github.com/Appliance-Surge/Appliance-Surge/internal/handlers"
	"github.com/Appliance-Surge/Appliance-Surge/internal/storage"
)

// Initializes the routing for the web server
//
// Params:
// - cfg (*config.Config): The configuration instance
// - db (*storage.DB): The database instance used by the handlers
//
// Returns:
// - *chi.Mux: A new chi router instance.
//
// Since: 0.0.1
func NewRouter(cfg *config.Config, db *storage.DB) *chi.Mux {
	r := chi.NewRouter()
	Routes(r, db)
	SetupFileServers(r)
	return r
}

// Sets the application routes
//
// Params:
// - r (*chi.Mux): Router to which the routes will be added
// - db (*storage.DB): The database instance used by the handlers
//
// Returns: void
//
// Since 0.0.1
func Routes(r *chi.Mux, db *storage.DB) {
	r.Get("/", handlers.HomeHandler(db))
}

// Serve asset files to client
//
// Params:
// - r (*hci.Mux): Router to which the file servers will be added.
//
// Returns: void
//
// Since 0.0.1
func SetupFileServers(r *chi.Mux) {
	staticFileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", staticFileServer))

	publicFileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public", publicFileServer))
}
