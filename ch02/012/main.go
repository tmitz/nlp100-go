package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]
	f, err := os.Open(arg)
	if err != nil {
		panic(err)
	}
	outCol(f)
}

func outCol(f *os.File) {
	f1, err := os.Create("col1.txt")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f2, err := os.Create("col2.txt")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	w1 := bufio.NewWriter(f1)
	w2 := bufio.NewWriter(f2)

	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		list := strings.Split(line, "\t")
		fmt.Fprintf(w1, "%s\n", list[0])
		fmt.Fprintf(w2, "%s\n", list[1])
	}
	w1.Flush()
	w2.Flush()
}
