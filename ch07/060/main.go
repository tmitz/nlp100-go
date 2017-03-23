package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"os"

	"time"

	"github.com/go-redis/redis"
)

type Result struct {
	Area     string  `json:"area,omitempty"`
	Name     string  `json:"name,omitempty"`
	SortName string  `json:"sort_name,omitempty"`
	Gid      string  `json:"gid,omitempty"`
	Type     string  `json:"type,omitempty"`
	ID       int     `json:"id,omitempty"`
	Begin    Begin   `json:"begin,omitempty"`
	Aliases  Aliases `json:"aliases,omitempty"`
	Tags     Tags    `json:"tags,omitempty"`
	Rating   Rating  `json:"rating,omitempty"`
}

type Alias struct {
	Name     string `json:"name"`
	SortName string `json:"sort_name"`
}

type Aliases []Alias

type Begin struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Date  int `json:"date"`
}

type Tag struct {
	Count int    `json:"count"`
	Value string `json:"value"`
}

type Tags []Tag

type Rating struct {
	Count int `json:"count"`
	Value int `json:"value"`
}

func main() {
	file := os.Args[1:][0]
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close() // nolint

	r, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer r.Close() // nolint

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		var data Result
		b := sc.Bytes()
		err := json.Unmarshal(b, &data)
		if err != nil {
			panic(err)
		}
		if data.Name != "" && data.Area != "" {
			err := client.Set(data.Name, data.Area, time.Hour).Err()
			if err != nil {
				panic(err)
			}
		}
	}
}
