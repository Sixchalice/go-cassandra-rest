package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func InitLogger(logPath string) (func(), error) {
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	cleanup := func() {
		file.Close()
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	handler := slog.NewJSONHandler(multiWriter, nil)
	logger := slog.New(handler)

	slog.SetDefault(logger)

	return cleanup, nil
}
