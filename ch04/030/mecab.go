package mecab

import (
	"bufio"
	"os"
	"strings"
)

type Mecab map[string]string
type Sentence []Mecab

func Load(file string) Sentence {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var out Sentence
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		text := sc.Text()
		if strings.Contains(text, "EOS") {
			continue
		}

		mecab := mapping(text)
		out = append(out, mecab)
	}

	return out
}

func mapping(text string) Mecab {
	slice := strings.Split(text, "\t")
	background := strings.Split(slice[1], ",")
	mecab := make(Mecab)
	mecab["surface"] = slice[0]
	mecab["base"] = background[6]
	mecab["pos"] = background[0]
	mecab["pos1"] = background[1]

	return mecab
}
