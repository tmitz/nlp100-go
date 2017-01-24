package cipher

import (
	"fmt"
	"testing"
)

func TestCipher(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{
			"Hi He Lied Because Boron Could Not Oxidize Fluorine.",
			"Hr Hv Lrvw Bvxzfhv Blilm Clfow Nlg Ocrwrav Foflirmv.",
		},
		{
			"Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.",
			"Nld I mvvw z wirmp, zoxlslorx lu xlfihv, zugvi gsv svzeb ovxgfivh rmeloermt jfzmgfn nvxszmrxh.",
		},
		{
			"Hr Hv Lrvw Bvxzfhv Blilm Clfow Nlg Ocrwrav Foflirmv.",
			"Hi He Lied Because Boron Could Not Oxidize Fluorine.",
		},
		{
			"Nld I mvvw z wirmp, zoxlslorx lu xlfihv, zugvi gsv svzeb ovxgfivh rmeloermt jfzmgfn nvxszmrxh.",
			"Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s to %s", test.str, test.want), func(t *testing.T) {
			if got := Cipher(test.str); got != test.want {
				t.Errorf("got: %s, want: %s", got, test.want)
			}
		})
	}
}
