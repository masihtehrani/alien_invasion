package usecases

import (
	"context"

	"github.com/masihtehrani/alien_invasion/internal/adaptors/logger"
	"github.com/masihtehrani/alien_invasion/internal/entities/interfaces"
	"github.com/masihtehrani/alien_invasion/internal/entities/structs"
)

const AlienFightCondition = 2

type UseCases struct {
	config config

	logger *logger.Logger

	dataStore interfaces.IDataStore

	worlds map[string]*structs.City
	aliens map[string]structs.Alien
}

type config struct {
	roundCount int
	alienCount int
}

func New(_ context.Context, roundCount int, alienCount int, logger *logger.Logger,
	dataStore interfaces.IDataStore) interfaces.IUseCases {
	return &UseCases{
		config: config{
			roundCount: roundCount,
			alienCount: alienCount,
		},
		logger:    logger,
		dataStore: dataStore,
		worlds:    make(map[string]*structs.City),
		aliens:    make(map[string]structs.Alien),
	}
}
