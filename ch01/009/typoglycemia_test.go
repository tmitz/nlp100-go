package typoglycemia

import "testing"

var str = "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

func TestTypoglycemia(t *testing.T) {
	TypoGlycemia(str)
}
