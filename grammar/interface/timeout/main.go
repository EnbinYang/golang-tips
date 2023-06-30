package main

import (
	"net/http"
	"time"
)

func readDB() string {
	time.Sleep(200 * time.Millisecond) // service
	return "success\n"
}

func readDBInterface(w http.ResponseWriter, req *http.Request) {
	var rsp string

	done := make(chan struct{}, 1)
	go func() {
		rsp = readDB()
		done <- struct{}{} // write data after job done
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		rsp = "timeout\n" // timeout: 100ms
	}

	w.Write([]byte(rsp))
}

func main() {
	http.HandleFunc("/", readDBInterface)      // root path
	http.ListenAndServe("localhost:7788", nil) // bind IP and Port
}
