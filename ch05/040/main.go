package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Morph struct {
	Surface, Base, Pos, Pos1 string
}

type Morphs []Morph

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	ms := make(Morphs, 0)
	for sc.Scan() {
		text := sc.Text()
		if strings.HasPrefix(text, "*") || strings.Contains(text, "EOS") {
			continue
		}
		line := strings.Split(text, "\t")
		parse := strings.Split(line[1], ",")
		ms = append(ms, Morph{Surface: line[0], Base: parse[6], Pos: parse[0], Pos1: parse[1]})
	}
	fmt.Println(ms[2])
}
