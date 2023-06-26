package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Student struct {
	Name   string
	Age    int
	Height float32
}

type Request struct {
	StudentId string `json:"student_id"`
}

func GetStudentInfo(id string) Student {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1",
		DB:       0,
	})
	ctx := context.Background()

	stu := Student{}
	for field, value := range client.HGetAll(ctx, id).Val() {
		if field == "Name" {
			stu.Name = value
		} else if field == "Age" {
			if age, err := strconv.Atoi(value); err != nil {
				log.Println("convert failed:", err)
			} else {
				stu.Age = age
			}
		} else if field == "Height" {
			if height, err := strconv.ParseFloat(value, 10); err != nil {
				log.Println("convert failed:", err)
			} else {
				stu.Height = float32(height)
			}
		}
	}

	return stu
}

// GET interface: get parameters from the front end
func GetName(ctx *gin.Context) {
	param := ctx.Query("student_id") // GET
	if len(param) == 0 {
		ctx.String(http.StatusBadRequest, "no student_id recv")
		return
	}

	stu := GetStudentInfo(param)
	ctx.String(http.StatusOK, stu.Name)
}

// POST interface: get parameters from the front end
func GetAge(ctx *gin.Context) {
	param := ctx.PostForm("student_id")
	if len(param) == 0 {
		ctx.String(http.StatusBadRequest, "no student_id recv")
		return
	}

	stu := GetStudentInfo(param)
	ctx.String(http.StatusOK, strconv.Itoa(stu.Age))
}

func GetHeight(ctx *gin.Context) {
	var param Request
	if err := ctx.BindJSON(&param); err != nil {
		ctx.String(http.StatusBadRequest, "student_id need be JSON format")
		return
	}

	stu := GetStudentInfo(param.StudentId)
	ctx.JSON(http.StatusOK, stu)
}

func main() {
	engine := gin.Default()
	engine.GET("/get_name", GetName) // set up routing
	engine.POST("/get_age", GetAge)
	engine.POST("/get_height", GetHeight)
	if err := engine.Run("localhost:6677"); err != nil {
		panic(err)
	} // bind
}
