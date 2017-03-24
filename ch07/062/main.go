package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	keys, err := client.Keys("*").Result()
	if err != nil {
		panic(err)
	}
	count := 0
	for _, key := range keys {
		val, err := client.Get(key).Result()
		if err != nil {
			panic(err)
		}
		if val == "Japan" {
			count++
		}
	}
	fmt.Printf("Japan area artist counts: %d\n", count)
}
