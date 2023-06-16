package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Name string
}

var instance *Singleton
var once sync.Once

func getInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Name: "Singleton Instance"}
	})
	return instance
}

func main() {
	s1 := getInstance()
	s2 := getInstance()
	fmt.Println(s1 == s2)
}
