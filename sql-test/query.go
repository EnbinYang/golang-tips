package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/huandu/go-sqlbuilder"
)

func main() {
	// 建立数据库连接
	conn := "debian-sys-maint:P7g9fAYfTM4YhtoC@tcp(localhost:3306)/webserver"
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return
	}
	defer db.Close()

	// 构建 SQL 查询
	builder := sqlbuilder.NewSelectBuilder()
	builder.Select("username", "password").From("user")
	sql, args := builder.Build()

	// 执行查询
	rows, err := db.Query(sql, args...)
	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		var name, password string
		err := rows.Scan(&name, &password)
		if err != nil {
			fmt.Println("Failed to scan row: ", err)
			continue
		}
		fmt.Println(name, password)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Error in query result: ", err)
	}
}
