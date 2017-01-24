package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNgram(t *testing.T) {
	tests := []struct {
		sequence interface{}
		want     []string
	}{
		{
			"I am an NLPer",
			[]string{"I ", " a", "am", "m ", " a", "an", "n ", " N", "NL", "LP", "Pe", "er"},
		},
		{
			[]string{"I", "am", "an", "NLPer"},
			[]string{"I-am", "am-an", "an-NLPer"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v to %v", test.sequence, test.want), func(t *testing.T) {
			if got := Ngram(test.sequence, 2); !reflect.DeepEqual(got, test.want) {
				t.Errorf("got = %v, want = %v", got, test.want)
			}
		})
	}
}
