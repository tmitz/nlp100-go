package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var artists = []string{"The White Stripes", "Oasis", "Tahiti 80", "サカナクション"}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	mget, err := client.MGet(artists...).Result()
	if err != nil {
		panic(err)
	}
	for i, v := range mget {
		switch v.(type) {
		case string:
			fmt.Printf("%s / %s\n", artists[i], v)
		default:
			fmt.Println("value type is not string...")
		}
	}
}
