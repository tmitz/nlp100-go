package chunk

import (
	"bufio"
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
	Dst    int
	Srcs   int
}

func (c *Chunk) init() {
	*c = Chunk{}
}

type Sentence []Chunk // nolint

func (s *Sentence) Add(c Chunk) {
	if len(c.Morphs) > 0 {
		*s = append(*s, c)
	}
}

func (s *Sentence) init() {
	*s = nil
}

type Sentences []Sentence // nolint

func (ss *Sentences) Add(s Sentence) {
	if len(s) > 0 {
		*ss = append(*ss, s)
	}
}

// func addChunk(chunk Chunk, chunks *Chunks) Chunks {
// 	ch := *chunks
// 	if len(chunk.Morphs) > 0 {
// 		ch = append(ch, chunk)
// 	}
// 	return ch
// }

func List(file string) Sentences {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var chunk Chunk
	sentence := make(Sentence, 0)
	sentences := make(Sentences, 0)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		text := sc.Text()
		line := strings.Split(text, " ")

		switch line[0] {
		case "*":
			sentence.Add(chunk)
			dst, srcs := parseDstSrcs(line)
			chunk = Chunk{Dst: dst, Srcs: srcs}
		case "EOS":
			sentence.Add(chunk)
			sentences.Add(sentence)
			chunk.init()
			sentence.init()
		default:
			chunk.Morphs = append(chunk.Morphs, parseMorph(text))
		}
	}

	return sentences
}

func parseDstSrcs(line []string) (dst, srcs int) {
	dst, err := strconv.Atoi(strings.Replace(line[2], "D", "", 1))
	if err != nil {
		panic(err)
	}
	srcs, err = strconv.Atoi(line[1])
	if err != nil {
		panic(err)
	}
	return
}

func parseMorph(text string) Morph {
	sp := strings.Split(text, "\t")
	body := strings.Split(sp[1], ",")
	return Morph{Surface: sp[0], Base: body[6], Pos: body[0], Pos1: body[1]}
}
