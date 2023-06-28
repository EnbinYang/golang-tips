package main

import (
	"context"
	"log"
	"time"
)

func ctxWithCancel() {
	ctx, cancel := context.WithCancel(context.TODO())
	timeCancel := time.Now()

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	select {
	case <-ctx.Done():
		timeEnd := time.Now()
		log.Println(timeEnd.Sub(timeCancel).Milliseconds()) // 100
		err := ctx.Err()
		log.Println("Error:", err)
	}
}

func main() {
	ctxWithCancel()
}
