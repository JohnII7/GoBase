package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

// cas, compare and swap,

func add() {
	atomic.AddInt32(&i, 1)
}
func sub() {
	atomic.AddInt32(&i, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second)
	fmt.Printf("i: %v\n", i)
}
