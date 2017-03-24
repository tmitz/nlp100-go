package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/tmitz/nlp100-go/ch07/060/musicbrainz"
)

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
		var data musicbrainz.Result
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
