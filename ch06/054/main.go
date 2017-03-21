package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	WORD  = regexp.MustCompile(`<word>(\w+)</word>`)
	LEMMA = regexp.MustCompile(`<lemma>(\w+)</lemma>`)
	POS   = regexp.MustCompile(`<POS>(\w+)</POS>`)
)

type Tag struct {
	Word, Lemma, Pos string
}

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	tag := Tag{}
	for sc.Scan() {
		if tag.Word != "" && tag.Lemma != "" && tag.Pos != "" {
			fmt.Printf("%s\t%s\t%s\n", tag.Word, tag.Lemma, tag.Pos)
			tag = Tag{}
		}
		t := strings.TrimSpace(sc.Text())
		if WORD.MatchString(t) {
			tag.Word = WORD.FindStringSubmatch(t)[1]
		}
		if LEMMA.MatchString(t) {
			tag.Lemma = LEMMA.FindStringSubmatch(t)[1]
		}
		if POS.MatchString(t) {
			tag.Pos = POS.FindStringSubmatch(t)[1]
		}
	}
}
