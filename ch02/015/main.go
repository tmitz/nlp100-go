package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var tail = flag.Int("n", 10, "line number")

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
	len := len(list)
	for i := len - *tail; i < len; i++ {
		fmt.Println(list[i])
	}
}
