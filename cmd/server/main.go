package main

import (
	"log"

	"log/slog"

	"github.com/Sixchalice/go-cassandra-rest/internal/app"
)

func main() {
	a, cleanup, err := app.InitApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer cleanup()

	slog.Info("App is running", "config", a.Config)

	// TODO: start HTTP server, REST routes, etc.
}
