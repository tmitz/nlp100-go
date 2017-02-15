package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/tmitz/nlp100-go/ch03/020/british"
)

func main() {
	file := os.Args[1:]
	body := british.Parse(file[0])
	slice := strings.Split(body, "\n")

	re := regexp.MustCompile("Category:([^]|*]+)")
	for _, v := range slice {
		if strings.Contains(v, "Category:") {
			m := re.FindStringSubmatch(v)
			fmt.Println(m[1])
		}
	}
}
