package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetStudentInfo(t *testing.T) {
	id := "id:258"
	stu := GetStudentInfo(id)

	if len(stu.Name) == 0 {
		t.Fail()
	} else {
		log.Printf("%+v\n", stu)
	}
}

// http request
func TestGetName(t *testing.T) {
	if rsp, err := http.Get("http://localhost:6677/get_name?student_id=id:258"); err != nil {
		log.Println("Get Failed:", err)
		t.Fail()
	} else {
		defer rsp.Body.Close() // close body
		if bt, err := ioutil.ReadAll(rsp.Body); err != nil {
			log.Println("ReadAll Failed:", err)
			t.Fail()
		} else {
			log.Println(string(bt))
		}
	}
}

func TestGetAge(t *testing.T) {
	if rsp, err := http.PostForm("http://localhost:6677/get_age", url.Values{"student_id": []string{"id:369"}}); err != nil {
		log.Println("POST Failed:", err)
		t.Fail()
	} else {
		defer rsp.Body.Close() // close body
		if bt, err := ioutil.ReadAll(rsp.Body); err != nil {
			log.Println("ReadAll Failed:", err)
			t.Fail()
		} else {
			log.Println(string(bt))
		}
	}
}

func TestGetHeight(t *testing.T) {
	client := http.Client{}
	reader := strings.NewReader(`{"student_id":"id:258"}`)
	request, err := http.NewRequest("POST", "http://localhost:6677/get_height", reader)
	if err != nil {
		log.Println("Create NewRequest Failed:", err)
		return
	}
	request.Header.Add("Content-Type", "application/json") // type

	if rsp, err := client.Do(request); err != nil {
		log.Println("POST Failed:", err)
		t.Fail()
	} else {
		defer rsp.Body.Close() // close body
		if bt, err := ioutil.ReadAll(rsp.Body); err != nil {
			log.Println("ReadAll Failed:", err)
			t.Fail()
		} else {
			var stu Student
			if err := json.Unmarshal(bt, &stu); err != nil { // deserialization
				log.Println(err)
				t.Fail()
			}
			log.Println("full info:", stu)
			log.Printf("%+v\n", stu.Height)
		}
	}
}
