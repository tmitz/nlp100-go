package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tmitz/nlp100-go/ch05/041/chunk"
)

type Pair chunk.Sentence   // nolint
type Pairs []Pair          // nolint
type PairSentences []Pairs // nolint

func main() {
	file := os.Args[1:]
	sentences := chunk.List(file[0])
	ps := make(PairSentences, 0)
	for _, sentence := range sentences {
		ps = append(ps, createPairs(sentence))
	}
	for _, pairs := range ps {
		for _, pair := range pairs {
			res := make([]string, 0)
			for _, chunk := range pair {
				res = append(res, fmt.Sprintf("%s", chunk))
			}
			fmt.Println(strings.Join(res, "\t"))
		}
	}

}

func createPairs(sentence chunk.Sentence) Pairs {
	pairs := make(Pairs, 0)
	for _, chunk := range sentence {
		pair := make(Pair, 0, 2)
		if chunk.Dst != -1 {
			pair = append(pair, chunk, sentence[chunk.Dst])
			pairs = append(pairs, pair)
		}
	}
	return pairs
}
