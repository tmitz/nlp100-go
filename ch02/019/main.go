package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Prefs []string

type Result struct {
	Count int
	Pref  string
}

func (r Result) String() string {
	return fmt.Sprintf("%d %s\n", r.Count, r.Pref)
}

type Results []Result

func (r Results) Len() int {
	return len(r)
}

func (r Results) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Results) Less(i, j int) bool {
	return r[i].Pref < r[j].Pref
}

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	prefs := make(Prefs, 0)
	for sc.Scan() {
		slice := strings.Split(sc.Text(), "\t")
		prefs = append(prefs, slice[0])
	}

	prefCounts := make(map[string]int)
	for _, v := range prefs {
		prefCounts[v] += 1
	}

	res := make(Results, 0, len(prefCounts))
	for k, v := range prefCounts {
		result := Result{Pref: k, Count: v}
		res = append(res, result)
	}
	sort.Sort(Results(res))
	fmt.Println(res)
}
