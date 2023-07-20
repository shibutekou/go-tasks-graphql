package logger

import (
	"golang.org/x/exp/slog"
	"os"
)

const (
	levelDebug = "debug"
	levelInfo  = "info"
)

func New(level string) *slog.Logger {
	var log *slog.Logger

	switch level {
	case levelDebug:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case levelInfo:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
