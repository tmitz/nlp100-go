package main

import "testing"

func TestStrTemplate(t *testing.T) {
	tests := []struct {
		x, y, z interface{}
		want    string
	}{
		{
			12,
			"気温",
			22.4,
			"12時の気温は22.4",
		},
	}

	for _, test := range tests {
		if got := StrTemplate(test.x, test.y, test.z); got != test.want {
			t.Errorf("got: %s, want: %s", got, test.want)
		}
	}
}
