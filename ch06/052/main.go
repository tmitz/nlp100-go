package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/dchest/stemmer/porter2"
)

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	porter2 := porter2.Stemmer
	for sc.Scan() {
		t := sc.Text()
		if len(t) == 0 {
			continue
		}
		sp := strings.Split(t, " ")
		for k, v := range sp {
			v := fmt.Sprint(regexp.MustCompile(`\W`).ReplaceAllString(v, ""))
			fmt.Printf("%s\t%s\n", v, porter2.Stem(v))
			if len(sp) == k+1 {
				fmt.Println()
			}
		}
	}
}
