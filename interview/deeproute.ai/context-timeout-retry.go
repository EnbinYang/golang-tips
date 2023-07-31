package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		retryCount := 0
		maxRetry := 3

		for {
			select {
			case <-ctx.Done():
				fmt.Println("timeout: stop retry")
				return
			default:
				if retryCount < maxRetry {
					fmt.Println("retry time:", retryCount+1)
					retryCount++
					time.Sleep(1 * time.Second)
				} else {
					fmt.Println("timeout: stop retry")
					return
				}
			}
		}
	}()

	time.Sleep(10 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}
