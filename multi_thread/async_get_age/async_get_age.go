package main

import (
	"fmt"
	"time"
)

type CallbackFunc func(userId int, result string)

type Response struct {
	userId int
	result string
}

func getUserAgeAsync(userId int, callback CallbackFunc) {
	fmt.Printf("proceeding task {id = %d}\n", userId)
	time.Sleep(3 * time.Second)

	age := fmt.Sprintf("%d", userId)
	fmt.Printf("get task {id = %d} result suc\n", userId)

	callback(userId, age)
}

func handleCallbackResult(userId int, result string) {
	fmt.Printf("async done, user id = %d, age = %s\n", userId, result)
}

func main() {
	resultChan := make(chan Response)

	userIds := []int{20, 25, 30, 35}
	for _, userId := range userIds {
		go getUserAgeAsync(userId, func(userId int, result string) {
			rsp := Response{
				userId: userId,
				result: result,
			}
			resultChan <- rsp
		})
	}

	fmt.Println("proceeding other operation...")
	time.Sleep(5 * time.Second)

	for range userIds {
		rsp := <-resultChan
		handleCallbackResult(rsp.userId, rsp.result)
	}
}

