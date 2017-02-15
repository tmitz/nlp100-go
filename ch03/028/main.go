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
			res[m[1]] = removeMarkup(m[2])
		}
	}
	pp.Print(res)
}

func removeMarkup(s string) string {
	text := regexp.MustCompile(`'{2,5}`).ReplaceAllString(s, "")
	text = regexp.MustCompile(`\[{2}([^|\]]+?\|)*(.+?)\]{2}`).ReplaceAllString(text, "$2")
	text = regexp.MustCompile(`\[(.+?)(\s[^\]]+)\]`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`\{{2}.+?\|.+?\|(.+?)\}{2}`).ReplaceAllString(text, "$1")
	return text
}
