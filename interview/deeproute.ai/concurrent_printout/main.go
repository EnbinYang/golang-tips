package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan struct{}, 5)

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(num int) {
			<-ch
			fmt.Println(num)
			wg.Done()
		}(i)
	}

	close(ch)

	wg.Wait()
}
