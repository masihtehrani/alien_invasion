package main

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/masihtehrani/alien_invasion/internal/adaptors/logger"
	"github.com/masihtehrani/alien_invasion/internal/ports/datastore/filesystem"
	"github.com/masihtehrani/alien_invasion/internal/usecases"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	logger := logger.New(ctx, ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	go interrupt(ctx, cancelFunc, logger)

	commands, err := getCommands(ctx)
	if err != nil {
		logger.Error.Printf("getcommand error >> %s \n", err)
		os.Exit(1)
	}

	iDataStore := filesystem.New(ctx, commands.dataDir)

	iUseCases := usecases.New(ctx, commands.roundCount, commands.alienCount, &logger, iDataStore)

	err = iUseCases.Runner(ctx)
	if err != nil {
		logger.Error.Printf("Runner error >> %s \n", err)
		os.Exit(1)
	}
}
