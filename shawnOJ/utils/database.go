package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/shawnoj?charset=utf8&parseTime=True&loc=Local"
	var err error
	Db, err = sql.Open("mysql", dsn) // ✅ 这里使用全局变量 Db，而不是创建新的局部变量
	if err != nil {
		log.Fatalf("Database open error: %v", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	fmt.Println("Successfully connected to database!")
}
