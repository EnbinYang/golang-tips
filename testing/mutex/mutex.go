package main

import (
	"fmt"
	"sync"
)

var (
	sum   = 0
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func main() {
	wg.Add(5)

	go addToSum(1)
	go addToSum(2)
	go addToSum(3)
	go addToSum(4)
	go addToSum(5)

	wg.Wait()
	fmt.Println("sum:", sum)
}

func addToSum(num int) {
	mutex.Lock()
	sum += num
	mutex.Unlock()
	wg.Done()
}
