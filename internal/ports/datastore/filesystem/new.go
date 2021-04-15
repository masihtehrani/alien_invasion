package filesystem

import (
	"context"

	"github.com/masihtehrani/alien_invasion/internal/entities/interfaces"
)

type DataStore struct {
	dataDir string
}

// New filled config data related data store and return implementation data store.
func New(_ context.Context, dataDir string) interfaces.IDataStore {
	return &DataStore{dataDir: dataDir}
}
