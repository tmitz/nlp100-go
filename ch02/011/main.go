package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		log.Println("require filename!!!")
		return
	}
	for _, arg := range files {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			continue
		}
		str := strings.Replace(string(data), "\t", " ", -1)
		fw, err := os.Create(arg)
		if err != nil {
			panic(err)
		}
		defer fw.Close()

		w := bufio.NewWriter(fw)
		_, err = w.WriteString(str)
		if err != nil {
			panic(err)
		}
		w.Flush()
	}
}
