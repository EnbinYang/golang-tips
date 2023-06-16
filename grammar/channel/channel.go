package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)
	ch1 <- struct{}{}

	wg.Add(3)
	start := time.Now().Unix()
	go print("goroutine-1", ch1, ch2)
	go print("goroutine-2", ch2, ch3)
	go print("goroutine-3", ch3, ch1)
	wg.Wait()

	end := time.Now().Unix()
	fmt.Printf("duration: %d\n", end-start)
}

func print(g string, inchan chan struct{}, outchan chan struct{}) {
	time.Sleep(1 * time.Second)
	select {
	case <-inchan:
		fmt.Println(g)
		outchan <- struct{}{}
	}
	wg.Done()
}
