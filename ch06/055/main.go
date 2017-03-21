package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	WORD = regexp.MustCompile(`<word>(\w+)</word>`)
	NER  = regexp.MustCompile(`<NER>(\w+)</NER>`)
)

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var word string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := strings.TrimSpace(sc.Text())
		if WORD.MatchString(t) {
			word = WORD.FindStringSubmatch(t)[1]
		}

		if NER.MatchString(t) {
			if NER.FindStringSubmatch(t)[1] == "PERSON" {
				fmt.Println(word)
				word = ""
			}
		}
	}
}
