package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var email string
	fmt.Println("请输入name, age, email, 用空格分割")
	fmt.Scan(&name, &age, &email)
	if age > 28 {
		fmt.Println("太大了")
	} else {
		fmt.Println("OK")
	}
}
