package main

import "fmt"

const str = "stressed"

func main() {
	fmt.Println(Reverse(str))
}

func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
