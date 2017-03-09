package main

import (
	"fmt"
	"os"

	"github.com/tmitz/nlp100-go/ch05/040/morph"
)

func main() {
	file := os.Args[1:]
	res := morph.List(file[0])
	fmt.Println(res[2])
}
