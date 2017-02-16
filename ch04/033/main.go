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
	nouns := make([]string, 0)

	for _, v := range sentence {
		if v["pos"] == "名詞" && v["pos1"] == "サ変接続" {
			if _, ok := set[v["base"]]; !ok {
				set[v["base"]] = true
				nouns = append(nouns, v["base"])
			}
		}
	}

	fmt.Println(nouns)
}
