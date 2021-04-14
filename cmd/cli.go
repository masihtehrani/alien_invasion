package main

import (
	"context"
	"flag"

	"github.com/masihtehrani/alien_invasion/internal/adaptors/errors"
)

type commands struct {
	roundCount int
	alienCount int
	dataDir    string
}

// getCommands initialize cli and get command of user.
func getCommands(_ context.Context) (*commands, error) {
	roundCount := flag.Int("rounds", 10000, "count of round wander in the world")
	alienCount := flag.Int("aliens", 3, "count of aliens")
	dataDir := flag.String("data-dir", "testdata/world.txt", "dataDir read world data")
	flag.Parse()

	if roundCount == nil || *roundCount <= 0 {
		return nil, errors.ErrRoundCountEmpty
	} else if alienCount == nil || *alienCount <= 0 {
		return nil, errors.ErrAlienCountEmpty
	} else if dataDir == nil || *dataDir == "" {
		return nil, errors.ErrWorldFileEmpty
	}

	return &commands{
		alienCount: *alienCount,
		roundCount: *roundCount,
		dataDir:    *dataDir,
	}, nil
}
