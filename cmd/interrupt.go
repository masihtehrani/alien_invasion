package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/masihtehrani/alien_invasion/internal/adaptors/logger"
)

// interrupt handle kill signal in process software.
func interrupt(_ context.Context, cancelFunc context.CancelFunc, logger logger.Logger) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	defer close(c)

	for range c {
		cancelFunc()
		logger.Info.Printf("application stopped")

		return
	}
}
