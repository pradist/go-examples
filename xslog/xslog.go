package xslog

import (
	"log/slog"
	"os"
)

var programLevel = new(slog.LevelVar)

func Print() {

	logHandle := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     programLevel,
		AddSource: true,
	})

	logger := slog.New(logHandle)
	logger.Debug("Debug Level")
	logger.Info("Info Level")
	logger.Warn("Warn Level")
	logger.Error("Error Level")
}