package main

import (
	"fmt"
	"os"

	"github.com/tmitz/nlp100-go/ch04/030"
)

func main() {
	file := os.Args[1:]
	sentence := mecab.Load(file[0])
	set := make(map[string]bool)
	bases := make([]string, 0)

	for _, v := range sentence {
		if v["pos"] == "動詞" {
			w := v["base"]
			if _, ok := set[w]; !ok {
				set[w] = true
				bases = append(bases, w)
			}
		}
	}
	fmt.Println(bases)
}
