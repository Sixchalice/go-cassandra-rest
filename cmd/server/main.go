package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/Sixchalice/go-cassandra-rest/internal/app"
)

func main() {
	app, cleanup, err := app.InitApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer cleanup()

	slog.Info(fmt.Sprintf("App is running at http://localhost:%d/todo", app.Config.Port), "config", app.Config)
	http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Port), app.Mux)
}
