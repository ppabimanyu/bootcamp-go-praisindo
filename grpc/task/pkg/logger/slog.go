package logger

import (
	"io"
	"log/slog"
	"os"
)

type SlogConfig struct {
	LogPath string
	Debug   bool
}

func SetupLogger(config *SlogConfig) {
	var logger *slog.Logger
	err := os.MkdirAll(config.LogPath, 0755)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(config.LogPath+"/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	var logLevel slog.Level
	if config.Debug {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}
	logJSONHandler := slog.NewJSONHandler(io.MultiWriter(os.Stdout, file), &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	})
	logger = slog.New(logJSONHandler)
	slog.SetDefault(logger)
}
