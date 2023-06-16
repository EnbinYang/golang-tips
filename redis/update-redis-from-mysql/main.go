package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

type Cache struct {
	redisClient *redis.Client
	lock        sync.RWMutex
}

func NewCache() *Cache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1",
		DB:       0,
	})

	return &Cache{
		redisClient: redisClient,
		lock:        sync.RWMutex{},
	}
}

func fetchDataFromMySQL(key string) (string, error) {
	conn := "debian-sys-maint:P7g9fAYfTM4YhtoC@tcp(localhost:3306)/webserver"
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("MySQL error:", err)
		return "", err
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT password from user WHERE username = '%s'", key)
	fmt.Println("Query:", query)
	row := db.QueryRow(query)

	var data string
	err = row.Scan(&data)
	if err != nil {
		fmt.Println("Fetch error:", err)
		return "", err
	}

	fmt.Printf("Fetch data from MySQL, key: %s, value: %s", key, data)
	return data, nil
}

func (c *Cache) getKeyFromRedis(ctx context.Context, key string) (string, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	value, err := c.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		value, err := fetchDataFromMySQL(key)
		if err != nil {
			return "", err
		}

		c.lock.Lock()
		err = c.redisClient.Set(ctx, key, value, 1).Err()
		c.lock.Unlock()
		if err != nil {
			return "", err
		}

		return value, nil
	} else if err != nil {
		return "", err
	}

	fmt.Printf("Fetch data from MySQL, key: %s, value: %s", key, value)
	return value, nil
}

func main() {
	ctx := context.Background()
	cache := NewCache()

	// 第一次获取值, 若缓存中 Key 不存在, 直接从 MySQL 获取并更新缓存
	value, err := cache.redisClient.Get(ctx, "enbin").Result()
	if err != nil {
		fmt.Println("Warning:", err)
	} else {
		fmt.Println("Value:", value)
	}

	// 第二次获取值, 此时缓存中 Key 不存在, 直接从 Redis 获取
	value, err = cache.redisClient.Get(ctx, "enbin").Result()
	if err != nil {
		fmt.Println("Warning:", err)
	} else {
		fmt.Println("Value:", value)
	}

	// 后台异步更新
	go func() {
		for {
			// 定时更新缓存
			time.Sleep(1 * time.Second)
			value, err := fetchDataFromMySQL("enbin")
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			cache.lock.Lock()
			err = cache.redisClient.Set(ctx, "enbin", value, 60).Err()
			cache.lock.Unlock()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Success set data into MySQL")
			}
		}
	}()

	time.Sleep(10 * time.Second)

	value, err = cache.redisClient.Get(ctx, "enbin").Result()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", value)
	}
}
