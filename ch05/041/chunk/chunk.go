package chunk

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
	Dst    int
	Srcs   int
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

func (c *Chunk) init() {
	*c = Chunk{}
}

func (c *Chunk) Pair(sentence Sentence) string {
	return fmt.Sprintf("%s\t%s", c, sentence[c.Dst])
}

func (c *Chunk) HasNoun() bool {
	for _, morph := range c.Morphs {
		if morph.Pos == "名詞" {
			return true
		}
	}
	return false
}

func (c *Chunk) HasVerb() bool {
	for _, morph := range c.Morphs {
		if morph.Pos == "動詞" {
			return true
		}
	}
	return false
}

func (c *Chunk) HasParticle() bool {
	for _, morph := range c.Morphs {
		if morph.Pos == "助詞" {
			return true
		}
	}
	return false
}

func (c *Chunk) LastParticle() Morph {
	morphs := make(Morphs, 0)
	for _, morph := range c.Morphs {
		if morph.Pos == "助詞" {
			morphs = append(morphs, morph)
		}
	}
	return morphs[len(morphs)-1]
}

func (c *Chunk) FirstVerb() Morph {
	morphs := make(Morphs, 0)
	for _, morph := range c.Morphs {
		if morph.Pos == "動詞" {
			morphs = append(morphs, morph)
		}
	}
	return morphs[0]
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
