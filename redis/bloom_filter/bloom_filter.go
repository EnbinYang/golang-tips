package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/willf/bloom"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1",
		DB:       0,
	})
	defer redisClient.Close()

	conn := "debian-sys-maint:P7g9fAYfTM4YhtoC@tcp(localhost:3306)/webserver"
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}

	// create bloom filter
	bf := bloom.NewWithEstimates(1000000, 0.01)

	// select data from MySQL
	rows, err := db.Raw("SELECT username, password FROM user").Rows()
	if err != nil {
		log.Fatal("Error executing MySQL query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var uname string
		var password string
		err := rows.Scan(&uname, &password)
		if err != nil {
			log.Fatal("Error scanning MySQL row:", err)
			continue
		}

		// insert data into Redis
		err = redisClient.Set(ctx, fmt.Sprintf("username:%s", uname), password, 0).Err()
		if err != nil {
			log.Fatal("Error setting Redis:", err)
			continue
		}

		// insert data into bloom filter
		bf.Add([]byte(fmt.Sprintf("%s", uname)))
	}

	queryName := "enbin"
	var password string
	if bf.Test([]byte(fmt.Sprintf("%s", queryName))) {
		// bloom filter hit
		password, err = redisClient.Get(ctx, fmt.Sprintf("username:%s", queryName)).Result()
		if err == redis.Nil {
			log.Println("Data not found in Redis, key:", queryName)
			// query data from MySQL
			err = db.Raw("SELECT password FROM user WHERE username = ?", queryName).Scan(&password).Error
			if err != nil {
				log.Fatal("Error querying MySQL:", err)
			} else {
				log.Println("Data found in MySQL:", password)
			}
		} else if err != nil {
			log.Fatal("Error querying Redis:", err)
		} else {
			log.Println("Data found in Redis:", password)
		}
	} else {
		// NOT FOUND
		log.Println("Data not found")
	}

	// close MySQL connection
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal("Error closing MySQL connection:", err)
	}
	dbSQL.Close()
}
