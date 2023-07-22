package main

import (
	"context"
	"fmt"
	"time"
)

type CallbackFunc func(userId int, result string)

func getUserAgeAsync(ctx context.Context, userId int, callback CallbackFunc) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // timeout is two seconds
	defer cancel()

	fmt.Printf("proceeding task {id = %d}\n", userId)
	time.Sleep(3 * time.Second) // operation execution time is three seconds

	select {
	case <-time.After(3 * time.Second):
		age := fmt.Sprintf("%d", userId)
		fmt.Printf("get task {id = %d} result suc\n", userId)
		callback(userId, age)

	case <-ctx.Done():
		if ctx.Err() != nil {
			fmt.Printf("error: get task {id = %d} timeout\n", userId)
		}
	}
}

func handleCallbackResult(userId int, result string) {
	fmt.Printf("async done, user id = %d, age = %s\n", userId, result)
}

func main() {
	ctx := context.Background()

	userIds := []int{20, 25, 30, 35}
	for _, userId := range userIds {
		go getUserAgeAsync(ctx, userId, func(userId int, result string) {
			handleCallbackResult(userId, result)
		})
	}

	fmt.Println("proceeding other operation...")
	time.Sleep(5 * time.Second)
}

