package main

import (
	"bufio"
	"fmt"
	"strings"
)

const s = "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."

func main() {
	fmt.Println(SymbolOfElement(s))
}

func SymbolOfElement(s string) map[string]int {
	lists := splitWord(s)
	res := make(map[string]int, len(lists))

	for i, v := range lists {
		switch i {
		case 0, 4, 5, 6, 7, 8, 14, 15, 18:
			res[v[0:1]] = i
		default:
			res[v[0:2]] = i
		}
	}

	return res
}

func splitWord(s string) []string {
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(bufio.ScanWords)
	res := []string{}
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res
}
