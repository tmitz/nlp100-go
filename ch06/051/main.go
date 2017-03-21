package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()
		if len(t) == 0 {
			continue
		}
		sp := strings.Split(t, " ")
		for k, v := range sp {
			fmt.Println(regexp.MustCompile(`\W`).ReplaceAllString(v, ""))
			if len(sp) == k+1 {
				fmt.Println()
			}
		}
	}
}
