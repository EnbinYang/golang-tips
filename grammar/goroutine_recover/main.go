package main

import (
	"log"
	"time"
)

func foo() {
	defer func() {
		err := recover() // catch panic and print log
		if err != nil {
			log.Println("panic:", err) // logging
		}
	}()
	a, b := 3, 0
	log.Println(a, b)
	_ = a / b // panic
	log.Println("foo finish")
}

func main() {
	go foo()
	time.Sleep(1 * time.Second)
	log.Println("main finish")
}
