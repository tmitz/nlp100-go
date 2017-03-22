package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tmitz/nlp100-go/ch06/056/representative"
)

type Word struct {
	Type      string
	Dependent string
}
type Words []Word

type Gov struct {
	Idx      int
	Governor string
}

func main() {
	file := os.Args[1:][0]
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	dict := make(map[Gov]Words, 0)
	doc := representative.Init(b)
	for _, sentence := range doc.Document.Sentences.Sentence {
		for _, depends := range sentence.Dependencies {
			if depends.Type == "collapsed-dependencies" {
				for _, dep := range depends.Dependency {
					gov := Gov{Idx: dep.Governor.Idx, Governor: dep.Governor.Value}
					if dep.Type == "nsubj" || dep.Type == "dobj" {
						dict[gov] = append(dict[gov], Word{Type: dep.Type, Dependent: dep.Dependent})
					}
				}
			}
		}
	}

	verbs := make([]Gov, 0)
	for gov, words := range dict {
		verbset := make(map[string]bool, 0)
		for _, word := range words {
			if !verbset[word.Type] {
				verbset[word.Type] = true
			}
		}
		if len(verbset) == 2 {
			verbs = append(verbs, gov)
		}
	}
	var res [][]string
	for _, verb := range verbs {
		var nsubj Words
		var dobj Words
		for _, d := range dict[verb] {
			if d.Type == "nsubj" {
				nsubj = append(nsubj, d)
			}
			if d.Type == "dobj" {
				dobj = append(dobj, d)
			}
		}
		var triple []string
		triple = append(triple, verb.Governor)
		for _, n := range nsubj {
			triple = append(triple, n.Dependent)
		}
		for _, d := range dobj {
			triple = append(triple, d.Dependent)
		}
		res = append(res, triple)
	}

	for _, dt := range res {
		fmt.Printf("%s\t%s\t%s\n", dt[1], dt[0], dt[2])
	}
}
