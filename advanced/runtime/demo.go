package main

import (
	"fmt"
	"runtime"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}

func main() {
	go show("goroutine running...")

	// 让出时间片给子协程, 主协程不会结束

	runtime.Gosched()

	fmt.Println("end")
}
