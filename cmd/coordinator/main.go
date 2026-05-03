package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/llm-d/llm-d-coordinator/pkg/coordinator"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	srv, err := coordinator.New(ctx, logger)
	if err != nil {
		logger.Error("failed to initialize coordinator", "error", err)
		return 1
	}

	if err := srv.Run(ctx); err != nil {
		logger.Error("coordinator exited with error", "error", err)
		return 1
	}

	return 0
}
