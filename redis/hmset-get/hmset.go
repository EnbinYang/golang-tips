package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// 创建 Redis 客户端
	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1",
		DB:       0,
	})

	data := map[string]interface{}{
		"task_id": "1",
		"status":  0,
		"ctime":   "2023-05-24 15:22:00",
	}

	// 插入数据
	err := cli.HMSet(context.Background(), "myhash", data).Err()
	if err != nil {
		fmt.Println("Failed to write data to Redis: ", err)
		return
	}

	fields := []string{"task_id", "status", "ctime"}

	values, err := cli.HMGet(context.Background(), "myhash", fields...).Result()
	if err != nil {
		fmt.Println("Failed to read data from Redis: ", err)
		return
	}

	for idx, field := range fields {
		value := values[idx]
		fmt.Printf("%s: %v\n", field, value)
	}
}
