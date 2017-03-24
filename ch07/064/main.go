package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"os"

	"github.com/tmitz/nlp100-go/ch07/060/musicbrainz"

	mgo "gopkg.in/mgo.v2"
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

	client, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	c := client.DB("musicbrainz").C("artists")

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		var data musicbrainz.Result
		err = json.Unmarshal(sc.Bytes(), &data)
		if err != nil {
			panic(err)
		}
		err = c.Insert(data)
		if err != nil {
			panic(err)
		}
	}

	err = c.EnsureIndexKey("name")
	if err != nil {
		panic(err)
	}
	err = c.EnsureIndexKey("aliases.name")
	if err != nil {
		panic(err)
	}
	err = c.EnsureIndexKey("tags.value")
	if err != nil {
		panic(err)
	}
	err = c.EnsureIndexKey("rating.value")
	if err != nil {
		panic(err)
	}
}
