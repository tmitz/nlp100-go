package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`<word>(\w+)</word>`)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := strings.TrimSpace(sc.Text())
		if re.MatchString(t) {
			fmt.Println(re.FindStringSubmatch(t)[1])
		}
	}
}
