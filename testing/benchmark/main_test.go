package main

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
)

type Student struct {
	Name   string
	Age    int
	Gender bool
}

type Class struct {
	Id      string
	Student []Student
}

var (
	s = Student{"enbin", 23, true}
	c = Class{
		Id:      "CS",
		Student: []Student{s, s, s},
	}
)

func BenchmarkJSON(b *testing.B) {
	// b.N is a hunge number
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(c)
		var data Class
		json.Unmarshal(bytes, &data)
	}
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := sonic.Marshal(c)
		var data Class
		sonic.Unmarshal(bytes, &data)
	}
}
