package main

import (
	"fmt"
	"time"
)

func main() {
	go count("goroutine 1", 5)
	go count("goroutine 2", 5)

	time.Sleep(6 * time.Second)

	fmt.Println("Exit")
}

func count(name string, n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(1 * time.Second)
	}
}
