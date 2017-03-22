package representative

import (
	"encoding/xml"
	"strings"
)

type Governor struct {
	Idx   int    `xml:"idx,attr"`
	Value string `xml:",chardata"`
}

type Dep struct {
	Type      string   `xml:"type,attr"`
	Governor  Governor `xml:"governor"`
	Dependent string   `xml:"dependent"`
}

type Dependency struct {
	Type       string `xml:"type,attr"`
	Dependency []Dep  `xml:"dep"`
}

type Token struct {
	Word string `xml:"word"`
}

type Tokens struct {
	Token []Token `xml:"token"`
}

type Sentence struct {
	Tokens       []Tokens     `xml:"tokens"`
	Dependencies []Dependency `xml:"dependencies"`
}

type Sentences struct {
	Sentence []Sentence `xml:"sentence"`
}

type Mention struct {
	Representative bool   `xml:"representative,attr"`
	Sentence       int    `xml:"sentence"`
	Start          int    `xml:"start"`
	End            int    `xml:"end"`
	Head           int    `xml:"head"`
	Text           string `xml:"text"`
}

type Coreference struct {
	Mention []Mention `xml:"mention"`
}

func (c *Coreference) represent() Mention {
	var res Mention
	for _, m := range c.Mention {
		if m.Representative {
			res = m
		}
	}
	return res
}

type Coreferences struct {
	Coreference []Coreference `xml:"coreference"`
}

type Document struct {
	Sentences    Sentences    `xml:"sentences"`
	Coreferences Coreferences `xml:"coreference"`
}

type Root struct {
	XMLName  xml.Name `xml:"root"`
	Document Document `xml:"document"`
}

func Init(b []byte) Root {
	data := Root{}
	err := xml.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func ListSentences(data Root) [][]string {
	var sentences [][]string
	for _, s := range data.Document.Sentences.Sentence {
		for _, ts := range s.Tokens {
			var t1 []string
			for _, t := range ts.Token {
				t1 = append(t1, t.Word)
			}
			sentences = append(sentences, t1)
		}
	}
	return sentences
}

func RepresentativeSentences(data Root) [][]string {
	sentences := ListSentences(data)
	for _, coref := range data.Document.Coreferences.Coreference {
		for _, m := range coref.Mention {
			if !m.Representative {
				sentence := m.Sentence - 1
				start := m.Start - 1
				end := m.End - 2

				target := sentences[sentence]
				target[start] = strings.TrimSpace(coref.represent().Text) + "(" + target[start]
				target[end] = target[end] + ")"
				sentences[sentence] = target
			}
		}
	}
	return sentences
}
