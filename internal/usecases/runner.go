package usecases

import (
	"context"
	"fmt"
)

// Runner it's just exposed method and call all logic with this method.
func (u *UseCases) Runner(ctx context.Context) error {
	textWorlds, err := u.dataStore.GetDataWorld(ctx)
	if err != nil {
		return fmt.Errorf("Runner >> u.dataStore.GetDataWorld >> %w", err)
	}

	u.createWorlds(ctx, textWorlds)
	u.createAliens(ctx)
	u.assignAliensToCities(ctx)

	for i := 0; i < u.config.roundCount; i++ {
		u.MoveAliens(ctx)
	}

	fmt.Println(u.printWorld())

	return nil
}
