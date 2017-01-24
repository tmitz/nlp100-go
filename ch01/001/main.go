package main

import (
	"bytes"
	"fmt"
)

const str = "パタトクカシーー"

func main() {
	fmt.Println(ConcatOdd(str))
}

func ConcatOdd(str string) string {
	var buf bytes.Buffer
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		if i%2 == 1 {
			continue
		}
		buf.WriteRune(r[i])
	}
	return buf.String()
}
