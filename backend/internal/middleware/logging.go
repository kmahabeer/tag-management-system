package middleware

import (
	"log/slog"
	"os"

	"github.com/kmahabeer/tag-management-system/backend/internal/config"
)

// InitLogger initializes the logger based on the provided configuration
func InitLogger(cfg *config.Config) error {
	var level slog.Level
	switch cfg.Logging.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{Level: level}

	var writer *os.File
	if cfg.Logging.Output == "stdout" {
		writer = os.Stdout
	} else {
		var err error
		writer, err = os.OpenFile(cfg.Logging.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
	}

	if cfg.Logging.Format == "json" {
		handler = slog.NewJSONHandler(writer, opts)
	} else {
		handler = slog.NewTextHandler(writer, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return nil
}
