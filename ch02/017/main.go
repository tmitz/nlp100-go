package main

import "os"
import "bufio"
import "fmt"
import "strings"
import "sort"

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var sl []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		slice := strings.Split(sc.Text(), "\t")
		sl = append(sl, slice[0])
	}
	res := uniqueStringSlice(sl)
	sort.Sort(sort.StringSlice(res))
	for _, v := range res {
		fmt.Println(v)
	}
}

func uniqueStringSlice(ss []string) []string {
	set := make(map[string]bool)
	res := make([]string, 0, len(ss))
	for _, v := range ss {
		if _, ok := set[v]; !ok {
			set[v] = true
			res = append(res, v)
		}
	}
	return res
}
