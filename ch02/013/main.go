package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]

	d1, err := ioutil.ReadFile(files[0])
	if err != nil {
		panic(err)
	}
	d2, err := ioutil.ReadFile(files[1])
	if err != nil {
		panic(err)
	}
	a1 := strings.Split(string(d1), "\n")
	a2 := strings.Split(string(d2), "\n")

	for i := 0; i < len(a1); i++ {
		fmt.Printf("%s\t%s\n", a1[i], a2[i])
	}
}
