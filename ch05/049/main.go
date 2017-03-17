package main

import (
	"fmt"
	"os"

	"github.com/tmitz/nlp100-go/ch05/041/chunk"
)

type NounPair []chunk.Chunk
type NounPairs []NounPair

func main() {
	file := os.Args[1:]
	sentences := chunk.List(file[0])

	for _, sentence := range sentences {
		npairs := nounPairs(sentence)
		if len(npairs) == 0 {
			continue
		}

		var pathciRoot chunk.Sentence
		// var pathcjRoot chunk.Sentence
		for _, n := range npairs {
			ci, cj := n[0], n[1]
			ci.ReplaceNoun("X")
			cj.ReplaceNoun("Y")

			pathciRoot = pathToRoot(ci, sentence)
			pathcjRoot = pathToRoot(cj, sentence)

			fmt.Println(cj)
			fmt.Println(pathciRoot)
			if containNounChunk(cj, pathciRoot) {
				fmt.Println("unko")
			} else {
				// fmt.Println(pathcjRoot)
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

func nounPairs(sentence chunk.Sentence) NounPairs {
	nounPairs := make(NounPairs, 0)
	var chunks []chunk.Chunk

	for _, c := range sentence {
		if c.HasNoun() {
			chunks = append(chunks, c)
		}
	}
	for i, v := range chunks {
		for _, vv := range chunks[i+1:] {
			nounPair := make(NounPair, 0, 2)
			nounPairs = append(nounPairs, append(nounPair, v, vv))
		}
	}

	return nounPairs
}

func containNounChunk(c chunk.Chunk, s chunk.Sentence) bool {
	cstr := fmt.Sprint(c)
	for _, v := range s {
		vstr := fmt.Sprint(v)
		if vstr == cstr {
			return true
		}
	}
	return false
}
