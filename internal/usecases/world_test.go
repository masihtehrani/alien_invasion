package usecases

import (
	"context"
	"testing"

	"github.com/masihtehrani/alien_invasion/internal/entities/structs"
)

func TestUseCases_createWorlds(t *testing.T) {
	tests := []struct {
		name       string
		textWorlds []string
		worlds     map[string]*structs.City
	}{
		{name: "empty", textWorlds: []string{}, worlds: nil},
		{name: "filled", textWorlds: []string{"Foo north=Bar west=Baz south=Qu-ux", "Bar south=Foo west=Bee"},
			worlds: map[string]*structs.City{
				"bar":   &structs.City{},
				"baz":   &structs.City{},
				"bee":   &structs.City{},
				"foo":   &structs.City{},
				"qu-ux": &structs.City{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			u := &UseCases{
				worlds: map[string]*structs.City{},
			}
			u.createWorlds(ctx, tt.textWorlds)
			switch {
			case u.worlds == nil:
				if tt.worlds != nil {
					t.Errorf("world must be nil")
				}
			case len(u.worlds) > 0:
				if len(u.worlds) != len(tt.worlds) {
					t.Errorf("two worlds not equal len!")
				}
				for name := range u.worlds {
					if _, ok := tt.worlds[name]; !ok {
						t.Errorf("two worlds not equal!")
					}
				}
			}
		})
	}
}
