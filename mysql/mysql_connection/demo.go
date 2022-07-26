package main

import (
	"database/sql"
	"fmt"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:root@tcp(47.106.181.104:3306)/go_db"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试与数据库建立连接(检验dsn是否正确)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insert() {
	s := "insert into user(username,password)values (?,?)"
	r, err := db.Exec(s, "John117", "root")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
}

func insertByParams(username string, password string) {
	s := "insert into user(username,password)values (?,?)"
	r, err := db.Exec(s, username, password)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
}

func main() {
	/* db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	// fmt.Printf("db: %v\n", db) */

	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	// insert()
	insertByParams("郭宏志", "114514")
}
