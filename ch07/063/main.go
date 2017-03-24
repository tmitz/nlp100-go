package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"

	"strconv"

	"github.com/go-redis/redis"
	"github.com/tmitz/nlp100-go/ch07/060/musicbrainz"
)

var artists = []string{"Oasis", "Blur"}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	file := os.Args[1:][0]
	buildTagDB(client, file)
	printArtistTags(client, artists)
}

func buildTagDB(client *redis.Client, file string) {
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

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		var data musicbrainz.Result
		b := sc.Bytes()
		err = json.Unmarshal(b, &data)
		if err != nil {
			panic(err)
		}
		if data.Name != "" && data.Tags != nil {
			for _, tag := range data.Tags {
				err := client.HSet(data.Name, tag.Value, tag.Count).Err()
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func printArtistTags(client *redis.Client, artists []string) {
	for _, artist := range artists {
		hkeys, err := client.HKeys(artist).Result()
		if err != nil {
			panic(err)
		}
		counts, err := client.HMGet(artist, hkeys...).Result()
		if err != nil {
			panic(err)
		}

		res := make(map[string]int, 0)
		for i, hkey := range hkeys {
			if count, ok := counts[i].(string); ok {
				n, err := strconv.Atoi(count)
				if err != nil {
					panic(err)
				}
				res[hkey] = n
			}
		}

		fmt.Printf("--- %s tags ---\n", artist)
		for k, v := range res {
			fmt.Printf("tag: %s, count: %d\n", k, v)
		}
		fmt.Println()
	}
}
