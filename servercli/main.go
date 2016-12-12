package main

import (
	"flag"
	"github.com/dyeduguru/memkv/server"
)

var (
	port = flag.Int("port", 5555, "Port to use for server")
	host = flag.String("host", "localhost", "Host on which server is listening")
)

func main() {
	flag.Parse()
	cfg := &server.ServerConfig{Host: *host, Port: *port}
	server := server.Server(cfg)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
