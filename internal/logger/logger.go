package logger

import (
	"log/slog"
	"os"
)

func Logger(lev slog.Level) *slog.LevelVar {
	var programLevel = new(slog.LevelVar) // Info by default
	programLevel.Set(lev)
	l := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel}))
	slog.SetDefault(l)
	return programLevel
}
