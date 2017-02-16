package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/tmitz/nlp100-go/ch04/030"
)

type Word struct {
	Surface string
	Count   int
}

type Words []Word

func (w Words) Len() int {
	return len(w)
}

func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w Words) Less(i, j int) bool {
	return w[i].Count > w[j].Count
}

func main() {
	file := os.Args[1:]
	sentence := mecab.Load(file[0])
	freq := make(map[string]int)
	for _, v := range sentence {
		freq[v["base"]] += 1
	}

	ws := Words{}
	for k, v := range freq {
		w := Word{Surface: k, Count: v}
		ws = append(ws, w)
	}

	sort.Sort(ws)
	fmt.Println(ws)
}
