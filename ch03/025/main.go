package main

import (
	"fmt"
	"os"
	"strings"

	"regexp"

	"github.com/tmitz/nlp100-go/ch03/020/british"
)

func main() {
	file := os.Args[1:]
	body := british.Parse(file[0])

	re := regexp.MustCompile(`(?s){{基礎情報(.*?)[^a-z]}}`)
	basic := re.FindString(body)

	slice := strings.Split(basic, "\n")
	re2 := regexp.MustCompile(`\|(.+)`)
	res := make(map[string]string)
	for _, v := range slice {
		if !strings.Contains(v, "=") {
			continue
		}

		if kv := re2.FindStringSubmatch(v); len(kv) > 0 {
			sp := strings.Split(kv[1], " = ")
			res[sp[0]] = sp[1]
		}
	}

	fmt.Printf("%q\n", res)
}
