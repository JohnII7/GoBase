package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go showMsg("Hello, World===========")
	showMsg("goruntine")
}
