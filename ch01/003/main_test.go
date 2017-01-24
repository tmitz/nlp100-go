package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWordLengthList(t *testing.T) {
	tests := []struct {
		sentence string
		want     []int
	}{
		{
			"Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.",
			[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9},
		},
		{
			"OK, Google.",
			[]int{2, 6},
		},
		{
			"Hi, siri.",
			[]int{2, 4},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s to %v", test.sentence, test.want), func(t *testing.T) {
			if got := WordLengthList(test.sentence); !reflect.DeepEqual(got, test.want) {
				t.Errorf("got = %v, want = %v", got, test.want)
			}
		})
	}
}
