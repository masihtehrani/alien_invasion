package interfaces

import "context"

// IUseCases shamanic logic system.
type IUseCases interface {
	Runner(ctx context.Context) error
}
