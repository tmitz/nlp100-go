package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/tmitz/nlp100-go/ch05/041/chunk"
)

type NounPair []chunk.Chunk
type NounPairs []NounPair

func main() {
	file := os.Args[1:][0]
	sentences := chunk.List(file)

	for _, sentence := range sentences[:100] {
		npairs := nounPairs(sentence)
		if len(npairs) == 0 {
			continue
		}
		var formerRoot chunk.Sentence
		var latterRoot chunk.Sentence
		for _, pair := range npairs {
			former, latter := pair[0], pair[1]
			former.ReplaceNoun("X")
			latter.ReplaceNoun("Y")

			formerRoot = pathToRoot(former, sentence)
			latterRoot = pathToRoot(latter, sentence)

			latterX := latter
			latterX.ReplaceNoun("X")
			if containNounChunk(latterX, formerRoot) {
				idx := findRootIndex(latterX, formerRoot)
				fmt.Print(joinChunksArrow(formerRoot[:idx+1]))
			} else {
				common := commonChunk(formerRoot, latterRoot)
				if common.Dst == 0 && common.Srcs == 0 {
					continue
				}
				idxf := findRootIndex(common, formerRoot)
				idxl := findRootIndex(common, latterRoot)

				var path []string
				path = append(path, joinChunksArrow(formerRoot[:idxf]), joinChunksArrow(latterRoot[:idxl]), fmt.Sprint(common))
				fmt.Println(strings.Join(path, " | "))
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

func joinChunksArrow(chunks []chunk.Chunk) string {
	var res []string
	for _, c := range chunks {
		res = append(res, fmt.Sprint(c))
	}
	return strings.Join(res, " -> ")
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

func commonChunk(former chunk.Sentence, latter chunk.Sentence) chunk.Chunk {
	ch := chunk.Chunk{}
	rformer := reverseRoot(former)
	rlatter := reverseRoot(latter)

	for i, cf := range rformer {
		if i > len(rlatter)-1 {
			continue
		}
		cl := rlatter[i]
		if cf.Srcs != cl.Srcs {
			ch = rformer[i-1]
			break
		}
	}

	return ch
}

func findRootIndex(c chunk.Chunk, root chunk.Sentence) int {
	idx := -1
	for i, v := range root {
		if fmt.Sprint(v) == fmt.Sprint(c) {
			idx = i
		}
	}
	return idx
}

func reverseRoot(root chunk.Sentence) chunk.Sentence {
	var reverse chunk.Sentence
	for i := len(root) - 1; i > -1; i-- {
		reverse = append(reverse, root[i])
	}
	return reverse
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
