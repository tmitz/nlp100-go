package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Morph struct { // nolint
	Surface, Base, Pos, Pos1 string
}

type Morphs []Morph // nolint

type Chunk struct { // nolint
	Morphs Morphs
	dst    int
	srcs   int
}

func (c Chunk) String() string {
	strs := make([]string, 0)
	for _, m := range c.Morphs {
		if m.Pos == "記号" {
			continue
		}
		strs = append(strs, m.Surface)
	}
	return strings.Join(strs, "")
}

type Chunks []Chunk // nolint

type Pair []Chunk // nolint

type Pairs []Pair // nolint

type PairSentences []Pairs

func main() {
	file := os.Args[1:]
	sentences := morph(file[0])
	ps := make(PairSentences, 0)
	for _, sentence := range sentences {
		ps = append(ps, createPair(sentence))
	}
	res := make([]string, 0)
	for _, pairs := range ps {
		for _, pair := range pairs {
			for _, chunk := range pair {
				res = append(res, fmt.Sprintf("%s", chunk))
			}
			fmt.Println(strings.Join(res, "\t"))
			res = []string{}
		}
	}

}

func createPair(sentence Chunks) Pairs {
	pair := make(Pair, 0)
	pairs := make(Pairs, 0)
	for _, chunk := range sentence {
		if chunk.dst != -1 {
			pair = append(pair, chunk, sentence[chunk.dst])
			pairs = append(pairs, pair)
			pair = nil
		}
	}
	return pairs
}

func morph(file string) []Chunks {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var chunk Chunk
	chunks := make(Chunks, 0)
	sentences := make([]Chunks, 0)

	for sc.Scan() {
		text := sc.Text()
		line := strings.Split(text, " ")
		if line[0] == "*" {
			chunks = addChunk(chunk, &chunks)
			dst, err := strconv.Atoi(strings.Replace(line[2], "D", "", 1))
			if err != nil {
				panic(err)
			}
			srcs, err := strconv.Atoi(line[1])
			if err != nil {
				panic(err)
			}
			chunk = Chunk{dst: dst, srcs: srcs}
		} else if line[0] == "EOS" {
			chunks = addChunk(chunk, &chunks)
			sentences = append(sentences, chunks)
			chunk = Chunk{}
			chunks = nil
		} else {
			tab := strings.Split(text, "\t")
			comma := strings.Split(tab[1], ",")
			morph := Morph{Surface: tab[0], Base: comma[6], Pos: comma[0], Pos1: comma[1]}
			chunk.Morphs = append(chunk.Morphs, morph)
		}

	}

	return sentences
}

func addChunk(chunk Chunk, chunks *Chunks) Chunks {
	ch := *chunks
	if len(chunk.Morphs) > 0 {
		ch = append(ch, chunk)
	}
	return ch
}
