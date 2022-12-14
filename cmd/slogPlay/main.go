package main

import (
	"log"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout))
	jsonlogger := slog.New(slog.NewJSONHandler(os.Stdout))

	// output: time=2022-12-14T13:08:21.652Z level=INFO msg="Hello slog"
	logger.Info("Hello slog", "key", "value")

	// output: {"time":"2022-12-14T13:21:18.423355Z","level":"INFO","msg":"Hello slog","key":"value"}
	jsonlogger.Info("Hello slog", "key", "value")

	slog.SetDefault(logger)
	// slog.SetDefault(jsonlogger)

	// time=2022-12-14T13:29:11.819Z level=INFO msg=Warning!
	log.Printf("Warning!")

	// time=2022-12-14T13:41:21.768Z level=INFO msg="Warning!%!(EXTRA string=key, string=value)"
	log.Printf("Warning!", "key", "value")

	slog.Info("Slog info")
	slog.Warn("Slog info")
}
