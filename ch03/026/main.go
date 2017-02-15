package main

import (
	"os"

	"regexp"

	"github.com/k0kubun/pp"
	"github.com/tmitz/nlp100-go/ch03/020/british"
)

func main() {
	file := os.Args[1:]
	body := british.Parse(file[0])

	lines := regexp.MustCompile(`\n[\|}]`).Split(body, -1)
	res := make(map[string]string)

	for _, line := range lines {
		re := regexp.MustCompile(`(?s)^(.*?)\s=\s(.*)`)
		m := re.FindStringSubmatch(line)
		if len(m) > 0 {
			res[m[1]] = regexp.MustCompile(`'{2,5}`).ReplaceAllString(m[2], "")
		}
	}
	pp.Print(res)
}
