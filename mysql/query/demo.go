package main

import (
	"database/sql"
	"fmt"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	id       int
	username string
	password string
}

func queryOneRow() {
	s := "select * from user where id = ?"
	var u User
	err := db.QueryRow(s, 2).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("u: %v\n", u)
	}
}

func manyRow() {
	s := "select * from user"
	r, err := db.Query(s)
	var u User
	defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for r.Next() {
			r.Scan(&u.id, &u.username, &u.password)
			fmt.Printf("u: %v\n", u)
		}
	}
}

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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	queryOneRow()
	fmt.Println("==================")
	manyRow()
}
