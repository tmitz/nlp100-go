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
	verbs := make([]string, 0)

	for _, v := range sentence {
		if v["pos"] == "動詞" {
			verb := v["surface"]
			if _, ok := set[verb]; !ok {
				set[verb] = true
				verbs = append(verbs, verb)
			}
		}
	}
	fmt.Println(verbs)
}
