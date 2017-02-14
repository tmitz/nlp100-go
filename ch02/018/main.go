package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Temp struct {
	Pref        string
	City        string
	Temperature float64
	Date        string
}

func (t Temp) String() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\n", t.Pref, t.City, t.Temperature, t.Date)
}

type Temps []Temp

func (t Temps) Len() int {
	return len(t)
}

func (t Temps) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Temps) Less(i, j int) bool {
	return t[i].Temperature < t[j].Temperature
}

func main() {
	file := os.Args[1:]
	f, err := os.Open(file[0])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	temps := make(Temps, 0)
	for sc.Scan() {
		slice := strings.Split(sc.Text(), "\t")
		temp, err := strconv.ParseFloat(slice[2], 64)
		if err != nil {
			panic(err)
		}
		item := Temp{Pref: slice[0], City: slice[1], Temperature: temp, Date: slice[3]}
		temps = append(temps, item)
	}
	sort.Sort(Temps(temps))
	fmt.Println(temps)
}
