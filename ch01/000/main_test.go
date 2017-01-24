package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{"stressed", "desserts"},
		{"golang", "gnalog"},
		{"php", "php"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s to %s", test.str, test.want), func(t *testing.T) {
			if got := Reverse(test.str); got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
			}
		})
	}
}
