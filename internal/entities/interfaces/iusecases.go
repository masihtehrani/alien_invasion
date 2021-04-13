package interfaces

import "context"

type IUseCases interface {
	Runner(ctx context.Context) error
}
