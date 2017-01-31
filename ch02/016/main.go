package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var splitLine = flag.Int("l", 10, "line number")

func main() {
	flag.Parse()

	args := os.Args[1:]
	file := args[len(args)-1]

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	var list []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		list = append(list, sc.Text())
	}

	seq := 1
	for i := 0; i < len(list); i += *splitLine {
		f, err := os.Create(fmt.Sprintf("x%02d.txt", seq))
		if err != nil {
			panic(err)
		}
		wb := bufio.NewWriter(f)
		last := len(list)
		if i+*splitLine < len(list) {
			last = i + *splitLine
		}
		sublist := list[i:last]
		for _, v := range sublist {
			wb.WriteString(v)
			wb.WriteString("\n")
		}
		wb.Flush()
		seq++
	}

}
