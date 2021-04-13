package usecases

import (
	"context"
	"testing"
)

func Test_oppositeDirections(t *testing.T) {
	tests := []struct {
		direction string
		want      string
	}{
		{direction: "east", want: "west"},
		{direction: "west", want: "east"},
		{direction: "north", want: "south"},
		{direction: "south", want: "north"},
		{direction: "", want: ""},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.direction, func(t *testing.T) {
			t.Parallel()
			if got := oppositeDirections(ctx, tt.direction); got != tt.want {
				t.Errorf("oppositeDirections() = %v, want %v", got, tt.want)
			}
		})
	}
}
