package server

import (
    "net/http"

    "github.com/Appliance-Surge/Appliance-Surge/internal/config"
)

func NewServer (cfg *config.Config, handler http.Handler) *http.Server {
    return &http.Server {
        Addr: ":" + cfg.Port,
        Handler: handler,
    }
}
