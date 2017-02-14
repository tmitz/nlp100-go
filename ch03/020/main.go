package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Page struct {
	Title, Text string
}

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fz, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer fz.Close()

	var p Page
	reader := bufio.NewReaderSize(fz, 4096)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		dec := json.NewDecoder(strings.NewReader(line))
		dec.Decode(&p)
		if p.Title == "イギリス" {
			fmt.Println(p.Text)
			break
		}
	}
}
