package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/tmitz/nlp100-go/ch03/020/british"
)

func main() {
	file := os.Args[1:]
	body := british.Parse(file[0])
	slice := strings.Split(body, "\n")

	re := regexp.MustCompile("((?:File|ファイル):[^|]+)")
	for _, v := range slice {
		if m := re.FindString(v); len(m) > 0 {
			m = strings.Replace(m, " ", "_", -1)
			m = url.QueryEscape(m)
			fmt.Printf("http://ja.wikipedia.org/wiki/%s\n", m)
		}
	}
}
