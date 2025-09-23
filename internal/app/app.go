package app

import (
	"log/slog"

	"github.com/Sixchalice/go-cassandra-rest/config"
	"github.com/Sixchalice/go-cassandra-rest/internal/constants"
	"github.com/Sixchalice/go-cassandra-rest/internal/db"
	"github.com/Sixchalice/go-cassandra-rest/logger"
	"github.com/gocql/gocql"
)

// App holds references to the main components
type App struct {
	Config *config.Config
	DB     *gocql.Session
}

// InitApp initializes the app: config, logger, database
// Returns the App struct and a cleanup function
func InitApp() (*App, func(), error) {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, err
	}

	// Initialize logger
	cleanupLogger, err := logger.InitLogger(cfg.LogPath)
	if err != nil {
		return nil, nil, err
	}

	// Initialize database
	dbSession, dbCleanup, err := db.InitDB(cfg.DatabaseURI, constants.KeyspaceName)
	if err != nil {
		cleanupLogger()
		return nil, nil, err
	}

	// Aggregate cleanup functions
	cleanup := func() {
		dbCleanup()
		cleanupLogger()
	}

	slog.Info("App initialized successfully", "config", cfg)
	return &App{
		Config: cfg,
		DB:     dbSession,
	}, cleanup, nil
}
