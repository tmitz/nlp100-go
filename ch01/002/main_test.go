package main

import (
	"fmt"
	"testing"
)

func TestConcatString(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   string
	}{
		{"パトカー", "タクシー", "パタトクカシーー"},
		{"ペンペン", "アップル", "ペアンッペプンル"},
		{"ドロイド", "ピコ太郎", "ドピロコイ太ド郎"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s + %s = %s", test.s1, test.s2, test.want), func(t *testing.T) {
			if got, _ := ConcatString(test.s1, test.s2); got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
			}
		})
	}
}

func TestConcatStringError(t *testing.T) {
	tests := []struct {
		s1, s2 string
	}{
		{"ペン", "アップル"},
		{"ペコ", "ピコ太郎"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s + %s = error", test.s1, test.s2), func(t *testing.T) {
			if _, err := ConcatString(test.s1, test.s2); err == nil {
				t.Errorf("Required error: %s + %s", test.s1, test.s2)
			}
		})
	}
}
