package main

import (
	"reflect"
	"testing"
)

func TestSymbolOfElement(t *testing.T) {
	tests := []struct {
		s    string
		want map[string]int
	}{
		{
			"Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.",
			map[string]int{"K": 18, "Li": 2, "Ne": 9, "B": 4, "C": 5, "N": 6, "Si": 13, "Cl": 16, "H": 0, "He": 1, "Mi": 11, "P": 14, "S": 15, "Ar": 17, "Ca": 19, "O": 7, "Na": 10, "Al": 12, "Be": 3, "F": 8},
		},
	}
	for _, test := range tests {
		if got := SymbolOfElement(test.s); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got = %v, want = %v", got, test.want)
		}
	}
}
