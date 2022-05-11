package main

import (
	"os"

	"github.com/kolson4282/tdd-bible-api/graph/resolver"
	"github.com/kolson4282/tdd-bible-api/server"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server.RunServer(port, &resolver.Resolver{})
}
