package app

import (
	"log/slog"
	"os"

	"simple-bank/internal/config"
)

func NewLogger(env config.Environment) *slog.Logger {
	if env.IsProduction() {
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	} else {
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}
}
