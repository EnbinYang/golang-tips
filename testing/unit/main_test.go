package main

import (
	"encoding/json"
	"log"
	"testing"
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

func TestJSON(t *testing.T) {
	bytes, err := json.Marshal(c)
	if err != nil {
		t.Fail()
	}

	var data Class
	if err := json.Unmarshal(bytes, &data); err != nil {
		t.Fail()
	}
	if !(c.Id == data.Id && len(c.Student) == len(data.Student)) {
		t.Fail()
	}

	log.Println("Success")
}
