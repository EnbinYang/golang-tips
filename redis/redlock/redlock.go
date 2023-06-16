package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedLock struct {
	cli        *redis.Client
	key        string
	expiration time.Duration
}

func NewRedLock(cli *redis.Client, key string, expiration time.Duration) *RedLock {
	return &RedLock{
		cli:        cli,
		key:        key,
		expiration: expiration,
	}
}

func (lock *RedLock) Lock() (bool, error) {
	ctx := context.TODO()

	suc, err := lock.cli.SetNX(ctx, lock.key, "locked", time.Duration(lock.expiration)).Result()
	if err != nil {
		return false, err
	}

	return suc, nil
}

func (lock *RedLock) UnLock() error {
	ctx := context.TODO()

	_, err := lock.cli.Del(ctx, lock.key).Result()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1",
		DB:       0,
	})

	ticker := time.NewTicker(2 * time.Second)
	lock := NewRedLock(cli, "mylock", 5*time.Second)
	count := 0

	for range ticker.C {
		suc, err := lock.Lock()
		if err != nil {
			log.Fatalf("Failed to acquire lock: %v", err)
		}
		if suc {
			fmt.Println("Lock acquired")

			fmt.Println("Doing") // operation

			err := lock.UnLock()
			if err != nil {
				log.Fatalf("Failed to release lock: %v", err)
			}
			fmt.Println("Lock released")
		} else {
			log.Fatalf("Failed to acquire lock")
		}

		count++
		if count == 10 {
			break
		}
	}

	err := cli.Close()
	if err != nil {
		log.Fatalf("Failed to close Redis client: %v", err)
	}
}
