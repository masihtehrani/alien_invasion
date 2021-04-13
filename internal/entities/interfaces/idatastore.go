package interfaces

import (
	"context"
)

type IDataStore interface {
	GetDataWorld(ctx context.Context) ([]string, error)
}
