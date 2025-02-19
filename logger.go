package main

import (
	"log/slog"
	"os"
)

var handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
})

var logger = slog.New(handler)
