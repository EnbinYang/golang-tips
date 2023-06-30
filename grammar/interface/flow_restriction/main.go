package main

import (
	"log"
	"sync/atomic"
	"time"
)

var (
	concurrent      int32
	concurrentLimit = make(chan struct{}, 10) // concurrency is 10
)

func readDB() bool {
	atomic.AddInt32(&concurrent, 1)
	log.Println("readDB() concurrency:", atomic.LoadInt32(&concurrent))
	time.Sleep(200 * time.Millisecond) // read data from DataBase
	atomic.AddInt32(&concurrent, -1)
	return true
}

func handler() {
	concurrentLimit <- struct{}{} // flow restriction
	readDB()
	<-concurrentLimit
}

func main() {
	for i := 1; i <= 100; i++ {
		go handler()
	}
	time.Sleep(10 * time.Second)
}
