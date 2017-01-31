package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var head = flag.Int("n", 10, "line number")

func main() {
	flag.Parse()

	args := os.Args[1:]
	file := args[len(args)-1]

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)
	i := 1

	fmt.Println(*head)

	for sc.Scan() {
		if i > *head {
			break
		} else {
			i++
			fmt.Println(sc.Text())
		}
	}
}
