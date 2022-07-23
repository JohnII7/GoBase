package main

import (
	"fmt"
	"runtime"
	"time"
)

func show() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			// 退出当前协程
			runtime.Goexit()
		}
	}
}

func main() {
	go show()
	time.Sleep(time.Second)
	fmt.Println("end")
}
