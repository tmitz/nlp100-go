package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/tmitz/nlp100-go/ch06/056/representative"
)

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

	doc := representative.Init(b)
	var dotSentences []string
	for i, sentence := range doc.Document.Sentences.Sentence {
		for _, d := range sentence.Dependencies {
			if d.Type == "collapsed-dependencies" {
				dotSentences = append(dotSentences, dependToDot(i+1, d))
			}
		}
	}
	if len(os.Args) > 1 {
		idx, err := strconv.Atoi(os.Args[1:][1])
		if err != nil {
			panic(err)
		}
		fmt.Println(dotSentences[idx])
	} else {
		for _, dotSentence := range dotSentences {
			fmt.Println(dotSentence)
		}
	}
}

func dependToDot(i int, depend representative.Dependency) string {
	title := fmt.Sprintf("digraph sentence %d", i+1)
	header := "{ graph [rankdir = LR]"
	body := ""
	for _, dep := range depend.Dependency {
		body += fmt.Sprintf("\"%s\"->\"%s\" [label = \"%s\"]; ", dep.Governor, dep.Dependent, dep.Type)
	}
	return title + header + body + "}"
}
