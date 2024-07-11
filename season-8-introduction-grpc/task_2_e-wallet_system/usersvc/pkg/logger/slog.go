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

type Slog struct {
	*slog.Logger
}

func NewSlog(config *SlogConfig) Logger {
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

	logger.Info("Logger setup complete")

	return &Slog{logger}
}

func (s *Slog) Type() string {
	return "slog"
}

func (s *Slog) Default() interface{} {
	return s.Logger
}
