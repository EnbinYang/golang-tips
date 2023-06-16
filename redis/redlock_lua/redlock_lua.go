package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisMutex struct {
	client     *redis.Client
	lockKey    string
	expiration time.Duration
}

func NewRedisMutex(client *redis.Client, lockKey string, expiration time.Duration) *RedisMutex {
	return &RedisMutex{
		client:     client,
		lockKey:    lockKey,
		expiration: expiration,
	}
}

// 加锁
func (m *RedisMutex) acquireLock(lockValue string) bool {
	// 循环获取锁, 获取失败则睡眠 100ms 再重试
	for {
		// SETNX 尝试获取锁, 设置过期时间为 expiration
		result, err := m.client.SetNX(context.Background(), m.lockKey, lockValue, m.expiration).Result()
		if err != nil {
			return false
		}
		if result {
			return true
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// 解锁
func (m *RedisMutex) releaseLock(lockValue string) bool {
	// 使用 Lua 脚本删除锁 (保证原子性)
	luaScript := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`

	result, err := m.client.Eval(context.Background(), luaScript, []string{m.lockKey}, lockValue).Result()
	if err != nil || result == nil {
		return false
	}

	return result.(int64) == 1
}

// 工作线程
func worker(mutex *RedisMutex, lockValue string, wg *sync.WaitGroup) {
	defer wg.Done()

	// 获取锁
	if mutex.acquireLock(lockValue) {
		fmt.Println("Lock acquired by worker", lockValue)
		time.Sleep(5 * time.Second)
		fmt.Println("Worker", lockValue, "completed")

		// 释放锁
		if mutex.releaseLock(lockValue) {
			fmt.Println("Lock released by worker", lockValue)
		} else {
			fmt.Println("Failed to release lock by worker", lockValue)
		}
	} else {
		fmt.Println("Failed to acquire lock by worker", lockValue)
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1",
		DB:       0,
	})

	// 设置锁的键名和过期时间
	lockKey := "mylock"
	expiration := 10 * time.Second

	// 定义一个 Redis 分布式锁
	mutex := NewRedisMutex(client, lockKey, expiration)

	// 创建多个 Worker 并启动
	var wg sync.WaitGroup
	numWorkers := 5

	for i := 0; i < numWorkers; i++ {
		// 生成锁的唯一标识
		lockValue := fmt.Sprintf("worker-%d", i+1)
		wg.Add(1)
		go worker(mutex, lockValue, &wg) // 启动一个线程
	}

	// 等待所有 worker 执行完成
	wg.Wait()

	// 关闭客户端
	client.Close()
}
