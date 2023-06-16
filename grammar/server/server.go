package main

import (
	"log"
	"net"
	"time"
	"fmt"
	"strings"
	"bufio"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:7788")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
    	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
    	time.Sleep(delay)
    	fmt.Fprintln(c, "\t", shout)
   	time.Sleep(delay)
    	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
    	input := bufio.NewScanner(c)
    	for input.Scan() {
        	go echo(c, input.Text(), 1*time.Second)
   	}
   	c.Close()
}
