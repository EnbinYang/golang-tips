package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func main() {
	ch := make(chan struct{})

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go increment(ch)
	}

	wg.Wait()
	close(ch)

	fmt.Println(atomic.LoadInt64(&counter))
}

func increment(ch chan struct{}) {
	atomic.AddInt64(&counter, 1)
	wg.Done()
	<-ch
}
