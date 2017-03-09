package morph

import (
	"bufio"
	"os"
	"strings"
)

type Morph struct { // nolint
	Surface, Base, Pos, Pos1 string
}

type Morphs []Morph // nolint

func List(file string) Morphs {
	res := make(Morphs, 0)

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		text := sc.Text()
		line := strings.Split(text, " ")
		switch line[0] {
		case "*", "EOS":
			continue
		default:
			list := strings.Split(line[0], "\t")
			parse := strings.Split(list[1], ",")
			res = append(res, Morph{Surface: list[0], Base: parse[6], Pos: parse[0], Pos1: parse[1]})
		}
	}

	return res
}
