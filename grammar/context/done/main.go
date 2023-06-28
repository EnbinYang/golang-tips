package main

import (
	"context"
	"log"
	"time"
)

func funcWithTimeoutParent() {
	parent, cancel1 := context.WithTimeout(context.TODO(), 1000*time.Millisecond)
	defer cancel1()
	timeParent := time.Now()

	time.Sleep(500 * time.Millisecond) // wait 500ms after start child process

	child, cancel2 := context.WithTimeout(parent, 1000*time.Millisecond) // 499 = 1000 - 500
	defer cancel2()
	timeChilld := time.Now()

	select {
	case <-child.Done():
		err := child.Err()
		timeEnd := time.Now()
		log.Println(timeEnd.Sub(timeParent).Milliseconds(), timeEnd.Sub(timeChilld).Milliseconds()) // 1000, 499
		log.Println("Error:", err)
	}
}

func main() {
	funcWithTimeoutParent()
}
