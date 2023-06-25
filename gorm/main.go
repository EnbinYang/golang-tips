package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id      int
	Uid     int
	Keyword string `gorm:"column:keywords"` // key name mapping
	Degree  string
	Gender  string
	City    string
}

func (User) TableName() string {
	return "user" // table name mapping
}

func query(client *gorm.DB, city string) *User {
	var user User
	user.Uid = 10086
	err := client.Select("id, city").Where("city=?", city).Limit(1).First(&user).Error
	if err != nil {
		log.Println("Error querying data:", err)
	}
	return &user
}

func main() {
	conn := "debian-sys-maint:P7g9fAYfTM4YhtoC@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True"
	client, err := gorm.Open(mysql.Open(conn), nil)
	if err != nil {
		log.Println("Error connecting MySQL:", err)
	}

	// query
	user := query(client, "SH")
	if user != nil {
		log.Printf("%+v\n", *user) // "%+v" print struct data
	} else {
		log.Println("user is nil")
	}

	// insert
	newUser := User{
		Uid:     8848,
		Keyword: "clam down",
		Degree:  "B",
		Gender:  "M",
		City:    "BJ",
	}
	if err = client.Create(&newUser).Error; err != nil {
		log.Println("Error creating record:", err)
	}

	// update
	res := client.Model(&User{}).Where("uid=?", 8848).Update("city", "TW")
	if res.Error != nil {
		log.Println("Error updating record:", res.Error)
	}

	// update multiple column
	res = client.Model(&User{}).Where("uid=?", 8848).Updates(
		map[string]interface{}{"city": "TW", "gender": "F"},
	)
	if res.Error != nil {
		log.Println("Error updating multiple record:", res.Error)
	}

	// delete
	res = client.Where("uid=?", 95535).Delete(&User{})
	if res.Error != nil {
		log.Println("Error deleting record:", res.Error)
	}
}
