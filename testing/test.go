package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.RWMutex
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		go func(i int) {
			lock.Lock()
			m[i] = i
			lock.Unlock()
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			lock.RLock()
			fmt.Println(i, m[i])
			lock.RUnlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
}
