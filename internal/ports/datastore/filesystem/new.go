package filesystem

import (
	"context"

	"github.com/masihtehrani/alien_invasion/internal/entities/interfaces"
)

type DataStore struct {
	dataDir string
}

func New(_ context.Context, dataDir string) interfaces.IDataStore {
	return &DataStore{dataDir: dataDir}
}
