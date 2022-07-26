package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func update() {
	s := "update user set username = ?,password = ? where id = ?"
	r, err := db.Exec(s, "郭志宏的妈", "114514", 5)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("影响行数: %v\n", i)
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
	update()
}
