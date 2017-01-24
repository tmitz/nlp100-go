package main

import (
	"bytes"
	"fmt"
	"log"
)

const (
	s1 = "パトカー"
	s2 = "タクシー"
)

func main() {
	res, err := ConcatString(s1, s2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func ConcatString(s1 string, s2 string) (string, error) {
	var buf bytes.Buffer
	r1 := []rune(s1)
	r2 := []rune(s2)

	if len(r1) != len(r2) {
		return "", fmt.Errorf("word length is diffrent, s1 = %s, s2 = %s", s1, s2)
	}

	for i := 0; i < len(r1); i++ {
		buf.WriteRune(r1[i])
		buf.WriteRune(r2[i])
	}

	return buf.String(), nil
}
