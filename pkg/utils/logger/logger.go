package logger

import (
	"BackendDeveloperVK-testTask/pkg/utils"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var slogLogger *slog.Logger
	switch env {
	case envLocal:
		file, err := os.OpenFile("logfile.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		slogLogger = slog.New(
			slog.NewTextHandler(
				file,
				&slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		slogLogger = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		slogLogger = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	utils.LOG = slogLogger

	return slogLogger
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
