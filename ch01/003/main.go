package main

import (
	"bufio"
	"fmt"
	"strings"
)

const str = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

func main() {
	fmt.Println(WordLengthList(str))
}

func WordLengthList(s string) []int {
	res := []int{}
	s = clearStringNoise(s)
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		res = append(res, len(sc.Text()))
	}
	return res
}

func clearStringNoise(s string) string {
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	return s
}
