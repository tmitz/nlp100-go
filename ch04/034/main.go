package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/tmitz/nlp100-go/ch04/030"
)

func main() {
	file := os.Args[1:]
	sentence := mecab.Load(file[0])
	set := make(map[string]bool)
	ab := make([]string, 0)

	for i := 1; i < len(sentence)-1; i++ {
		if sentence[i]["surface"] == "の" {
			if sentence[i-1]["pos"] == "名詞" && sentence[i+1]["pos"] == "名詞" {
				s := fmt.Sprintf("%sの%s\n", sentence[i-1]["base"], sentence[i+1]["base"])
				if _, ok := set[s]; !ok {
					set[s] = true
					ab = append(ab, s)
				}
			}
		}
	}
	sort.Sort(sort.StringSlice(ab))
	fmt.Println(ab)
}
