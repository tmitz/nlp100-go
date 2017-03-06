package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

type Morph struct { // nolint
	Surface, Base, Pos, Pos1 string
}

type Morphs []Morph // nolint

type Chunk struct {
	Morphs Morphs
	dst    int
	srcs   int
}

type Chunks []Chunk // nolint

func main() {
	file := os.Args[1:]
	chunks := morph(file[0])
	pp.Print(chunks[7])
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
			if len(chunks) > 0 {
				sentences = append(sentences, chunks)
			}
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
