package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		log.Println("require filename.")
		return
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		defer f.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			continue
		}
		counts[arg] = Countlines(f)
	}

	for file, line := range counts {
		fmt.Printf("\t%d %s\n", line, file)
	}

}

func Countlines(f *os.File) int {
	count := 0
	input := bufio.NewScanner(f)
	for input.Scan() {
		count++
	}
	return count
}
