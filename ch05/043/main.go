package main

import (
	"fmt"
	"os"

	"github.com/tmitz/nlp100-go/ch05/041/chunk"
)

func main() {
	file := os.Args[1:]
	sentences := chunk.List(file[0])

	for _, sentence := range sentences {
		for _, c := range sentence {
			if c.HasNoun() && c.Dst > -1 && sentence[c.Dst].HasVerb() {
				fmt.Println(c.Pair(sentence))
			}
		}
	}
}
