package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tmitz/nlp100-go/ch05/041/chunk"
)

func main() {
	file := os.Args[1:]
	sentences := chunk.List(file[0])

	for _, sentence := range sentences {
		for _, c := range sentence {
			var res []string
			if c.HasNoun() {
				r := pathToRoot(c, sentence)
				for _, v := range r {
					res = append(res, fmt.Sprint(v))
				}
				fmt.Println(strings.Join(res, " -> "))
			}
		}
	}
}

func pathToRoot(c chunk.Chunk, s chunk.Sentence) chunk.Sentence {
	res := make(chunk.Sentence, 0)
	if c.Dst == -1 {
		res.Add(c)
	} else {
		res.Add(c)
		res = append(res, pathToRoot(s[c.Dst], s)...)
	}
	return res
}
