package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	LRB      = regexp.MustCompile(`-LRB- `)
	RRB      = regexp.MustCompile(` -RRB-`)
	NOTATION = regexp.MustCompile(` ([,\.:;])`)
	LDQ      = regexp.MustCompile("`` ")
	RDQ      = regexp.MustCompile(` ''`)
	SQ       = regexp.MustCompile(` '`)
	SQS      = regexp.MustCompile(` 's`)
)

type Token struct {
	Word string `xml:"word"`
}

type Tokens struct {
	Token []Token `xml:"token"`
}

type Sentence struct {
	Tokens []Tokens `xml:"tokens"`
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

func convertRepresetative(s string) string {
	s = LRB.ReplaceAllString(s, "(")
	s = RRB.ReplaceAllString(s, ")")
	s = LDQ.ReplaceAllString(s, "\"")
	s = RDQ.ReplaceAllString(s, "\"")
	s = SQS.ReplaceAllString(s, "'s")
	s = SQ.ReplaceAllString(s, "'")
	s = NOTATION.ReplaceAllString(s, "$1")
	return s
}

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := Root{}
	xmlstr, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(xmlstr, &data)
	if err != nil {
		panic(err)
	}
	sentences := representativeSentences(data)

	for _, sentence := range sentences {
		s := strings.Join(sentence, " ")
		s = convertRepresetative(s)
		fmt.Println(s)
	}
}

func listSentences(data Root) [][]string {
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

func representativeSentences(data Root) [][]string {
	sentences := listSentences(data)
	for _, coref := range data.Document.Coreferences.Coreference {
		for _, m := range coref.Mention {
			if !m.Representative {
				sentence_i := m.Sentence - 1
				start := m.Start - 1
				end := m.End - 2

				target := sentences[sentence_i]
				target[start] = strings.TrimSpace(coref.represent().Text) + "(" + target[start]
				target[end] = target[end] + ")"
				sentences[sentence_i] = target
			}
		}
	}
	return sentences
}
