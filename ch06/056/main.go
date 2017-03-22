package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/tmitz/nlp100-go/ch06/056/representative"
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

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	data := representative.Init(b)
	sentences := representative.RepresentativeSentences(data)

	for _, sentence := range sentences {
		s := strings.Join(sentence, " ")
		s = convertRepresetative(s)
		fmt.Println(s)
	}
}
