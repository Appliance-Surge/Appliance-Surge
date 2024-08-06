package server

import (
	"net/http"

	"github.com/Appliance-Surge/Appliance-Surge/internal/config"
)

// Create the server
//
// Params:
// - cfg (*config.Config) Server configurations
// - handler (http.Handler) base library http handler
//
// Returns:
// - *http.Server The server
//
// Since: 0.0.1
func NewServer(cfg *config.Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
	}
}
