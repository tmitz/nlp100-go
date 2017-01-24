package main

import (
	"fmt"
	"testing"
)

func TestConcatOdd(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{"パタトクカシーー", "パトカー"},
		{"PSPBAKPL", "PPAP"},
		{"トイラスントプリ", "トランプ"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s to %s", test.str, test.want), func(t *testing.T) {
			if got := ConcatOdd(test.str); got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
			}
		})
	}
}
