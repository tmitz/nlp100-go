package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`(?P<punt>[\.]) (?P<head>[A-Z])`)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()
		if len(t) == 0 {
			continue
		}
		ss := re.SubexpNames()
		r := fmt.Sprintf("${%s}\n${%s}", ss[1], ss[2])
		fmt.Println(re.ReplaceAllString(t, r))
	}
}
