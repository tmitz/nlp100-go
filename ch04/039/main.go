package main

// import (
// 	"github.com/sbinet/go-gnuplot"
// )

// func main() {
// 	fname := ""
// 	persist := false
// 	debug := true

// 	p, err := gnuplot.NewPlotter(fname, persist, debug)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer p.Close()

// 	p.PlotX([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "some data")
// 	p.CheckedCmd("set terminal png")
// 	p.CheckedCmd("set output '038.png'")
// 	p.CheckedCmd("replot")

// 	p.CheckedCmd("q")
// 	return
// }

import (
	"fmt"
	"os"
	"sort"

	"bufio"

	"github.com/tmitz/nlp100-go/ch04/030"
)

type Word struct {
	Surface string
	Count   int
}

type Words []Word

func (w Words) Len() int {
	return len(w)
}

func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w Words) Less(i, j int) bool {
	return w[i].Count < w[j].Count
}

func main() {
	file := os.Args[1:]
	sentence := mecab.Load(file[0])
	freq := make(map[string]int)
	for _, v := range sentence {
		freq[v["base"]] += 1
	}

	ws := Words{}
	for k, v := range freq {
		w := Word{Surface: k, Count: v}
		ws = append(ws, w)
	}

	sort.Sort(ws)

	hist := make(map[int]int)
	for _, v := range ws {
		hist[v.Count] += 1
	}

	var keys []int
	for k := range hist {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	outfile, err := os.Create("039.dat")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	buf := bufio.NewWriter(outfile)
	for _, k := range keys {
		s := fmt.Sprintf("%d %d\n", k, hist[k])
		buf.WriteString(s)
	}
	buf.Flush()
}
