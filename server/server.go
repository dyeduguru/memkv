package server

import (
	"fmt"
	"net/http"
)

type ServerConfig struct {
	Port int
	Host string
}

func Server(cfg *ServerConfig) *http.Server {
	str := NewStore()
	// Add handlers
	http.HandleFunc("/add-key", str.AddKey)
	http.HandleFunc("/get-value", str.GetValue)

	addr := fmt.Sprintf("%v:%d", cfg.Host, cfg.Port)
	return &http.Server{
		Addr: addr,
	}
}
