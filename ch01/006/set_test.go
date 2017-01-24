package set

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var (
	X = Ngram("paraparaparadise", 2)
	Y = Ngram("paragraph", 2)
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		x, y []string
		want []string
	}{
		{
			X, Y, []string{"ap", "ar", "pa", "ra"},
		},
	}

	for _, test := range tests {
		got := Intersection(X, Y)
		sort.Strings(got)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		x, y []string
		want []string
	}{
		{
			X, Y, []string{"ad", "ap", "ar", "di", "is", "pa", "ra", "se"},
		},
	}

	for _, test := range tests {
		got := Difference(X, Y)
		sort.Strings(got)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		x, y []string
		want []string
	}{
		{
			X, Y, []string{"ad", "ag", "ap", "ar", "di", "gr", "is", "pa", "ph", "ra", "se"},
		},
	}

	for _, test := range tests {
		got := Union(X, Y)
		sort.Strings(got)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
}

func TestIncludeBiGram(t *testing.T) {
	tests := []struct {
		bigram []string
		want   bool
	}{
		{X, true},
		{Y, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("IncludeGram(%v, \"se\") is %v", test.bigram, test.want), func(t *testing.T) {
			if got := IncludeGram(test.bigram, "se"); got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
