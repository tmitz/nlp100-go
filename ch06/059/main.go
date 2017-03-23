package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tmitz/nlp100-go/ch06/056/representative"
)

func main() {
	file := os.Args[1:][0]
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close() // nolint

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	doc := representative.Init(b)
	for _, sentence := range doc.Document.Sentences.Sentence {
		parse := strings.Replace(sentence.Parse, ")", " )", -1)
		parse = strings.Replace(parse, "(", "( ", -1)
		parses := strings.Split(parse, " ")

		for i, s := range parses {
			var list []string
			if s == "NP" {
				list = pickupNP(i, parses)
				printNP(list)
			}
		}
	}
}

func pickupNP(idx int, parses []string) []string {
	var res []string
	counter := 0
	for _, v := range parses[idx+1:] {
		if v == "(" {
			counter++
		}
		if v == ")" {
			counter--
		}
		if counter == -1 {
			fmt.Println()
			break
		}
		res = append(res, v)
	}
	return res
}

func printNP(list []string) {
	flag := 0
	for _, l := range list {
		if l == ")" {
			continue
		}
		if l == "(" {
			flag = 1
			continue
		}
		if l != "(" && flag == 1 {
			flag = -1
		} else {
			if l == "-LRB-" {
				fmt.Print("( ")
			} else if l == "-RRB-" {
				fmt.Print(") ")
			} else {
				fmt.Printf("%s ", l)
			}
		}
	}
}
