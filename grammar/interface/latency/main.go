package main

import (
	"log"
	"time"
)

func handler(a int, b int) string {
	timeStart := time.Now()
	// calculate interface latency using defer anonymous functions
	defer func() {
		log.Println("latency:", time.Since(timeStart).Milliseconds())
	}()

	if a > b {
		time.Sleep(100 * time.Millisecond)
		return "a"
	} else {
		time.Sleep(200 * time.Millisecond)
		return "b"
	}
}

func main() {
	handler(9, 7)
}
