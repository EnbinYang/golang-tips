package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func stringOperation(ctx context.Context, client *redis.Client) {
	key := "basic:enbin"
	value := "lol"

	// set
	if err := client.Set(ctx, key, value, 1*time.Second).Err(); err != nil {
		log.Println("Error setting Redis:", err)
	}

	// expire
	client.Expire(ctx, key, 3*time.Second)

	// get
	if value, err := client.Get(ctx, key).Result(); err != nil {
		log.Println("Error getting from Redis:", err)
	} else {
		log.Println("String value:", value)
	}

	// delete
	client.Del(ctx, key)
}

func listOperation(ctx context.Context, client *redis.Client) {
	key := "task_id"
	values := []interface{}{1, 2, 3}

	// RPush
	if err := client.RPush(ctx, key, values...).Err(); err != nil {
		log.Println("Error RPush to Redis:", err)
	}

	// LRange
	if value, err := client.LRange(ctx, key, 0, -1).Result(); err != nil {
		log.Println("Error LRange Redis:", err)
	} else {
		log.Println("LRange value:", value)
	}

	// delete
	client.Del(ctx, key)
}

func hashtableOperation(ctx context.Context, client *redis.Client) {
	// HSet
	if err := client.HSet(ctx, "id:258", "Name", "enbin", "Age", 23, "Height", 178).Err(); err != nil {
		log.Println("Error HSet Redis:", err)
	}
	if err := client.HSet(ctx, "id:369", "Name", "shaun", "Age", 99, "Height", 160).Err(); err != nil {
		log.Println("Error HSet Redis:", err)
	}

	// HGet
	if value, err := client.HGet(ctx, "id:258", "Age").Result(); err != nil {
		log.Println("Error HGet Redis:", err)
	} else {
		log.Println("HGet value:", value)
	}

	// HGetAll
	for field, value := range client.HGetAll(ctx, "id:369").Val() {
		log.Println(field, value)
	}

	// delete
	client.Del(ctx, "id:258") // key: id:258
	client.Del(ctx, "id:369")
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1",
		DB:       0,
	})
	ctx := context.Background() // null context

	stringOperation(ctx, client)
	listOperation(ctx, client)
	hashtableOperation(ctx, client)
}
