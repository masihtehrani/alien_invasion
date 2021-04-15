package usecases

import (
	"context"
	"fmt"
	"testing"

	"github.com/masihtehrani/alien_invasion/internal/entities/structs"
)

func Test_choseAlienName(t *testing.T) {
	tests := []struct {
		i    int
		want string
	}{
		{i: 0, want: "Liam_0"},
		{i: 5, want: "Ava_0"},
		{i: 9, want: "Isabella_0"},
		{i: 10, want: "Liam_1"},
		{i: 11, want: "Olivia_1"},
		{i: 19, want: "Isabella_1"},
		{i: 190, want: "Liam_19"},
		{i: 199, want: "Isabella_19"},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.i), func(t *testing.T) {
			t.Parallel()
			if got := choseAlienName(ctx, tt.i); got != tt.want {
				t.Errorf("choseAlienName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCases_createAliens(t *testing.T) {
	tests := []struct {
		name       string
		alienCount int
		alien      map[string]structs.Alien
	}{
		{name: "one", alienCount: 1, alien: map[string]structs.Alien{"Liam_0": {Name: "Liam_0", City: nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCases{
				config: config{alienCount: tt.alienCount},
				aliens: make(map[string]structs.Alien),
			}
			ctx := context.Background()
			u.createAliens(ctx)
			fmt.Println(u.aliens)
			fmt.Println(tt.alien)
			if !(fmt.Sprint(u.aliens) == fmt.Sprint(tt.alien)) {
				t.Errorf("aliens is not equal")
			}
		})
	}
}
