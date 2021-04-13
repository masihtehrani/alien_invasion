package usecases

import (
	"context"
	"fmt"
	"testing"
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
