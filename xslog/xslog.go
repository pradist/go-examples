package xslog

import (
	"log/slog"
	"os"
)

var programLevel = new(slog.LevelVar)

func createLogHandle() *slog.JSONHandler {
	return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     programLevel,
		AddSource: true,
		ReplaceAttr: func(groups []string, l slog.Attr) slog.Attr {
			if l.Key == slog.MessageKey {
				l.Key = "message"
				return l
			}
			return l
		},
	}).WithAttrs([]slog.Attr{
		slog.String("app", "xslog"),
		slog.String("version", "1.0.0"),
	}).(*slog.JSONHandler)
}

func Print() {
	logHandle := createLogHandle()

	logger := slog.New(logHandle)
	logger.Info("Hello World!")
	logger.Debug("Hello World!")
}
