package interfaces

import (
	"context"
)

// IDataStore shamanic data store system.
type IDataStore interface {
	GetDataWorld(ctx context.Context) ([]string, error)
}
