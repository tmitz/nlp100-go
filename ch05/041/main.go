package main

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/tmitz/nlp100-go/ch05/041/chunk"
)

func main() {
	file := os.Args[1:]
	sentences := chunk.List(file[0])
	pp.Print(sentences[7])
}
