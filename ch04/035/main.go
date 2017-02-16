package main

import (
	"fmt"
	"os"

	"github.com/tmitz/nlp100-go/ch04/030"
)

func main() {
	file := os.Args[1:]
	sentence := mecab.Load(file[0])

	seqNoun := make([]string, 0)
	seqNouns := make([][]string, 0)

	for i := 0; i < len(sentence); i++ {
		if sentence[i]["pos"] == "名詞" {
			seqNoun = append(seqNoun, sentence[i]["surface"])
		} else {
			if len(seqNoun) > 1 {
				seqNouns = append(seqNouns, seqNoun)
			}
			seqNoun = make([]string, 0)
		}
	}

	fmt.Println(seqNouns)
}
