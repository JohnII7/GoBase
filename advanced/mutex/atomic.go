package main

import (
	"fmt"
	"sync/atomic"
)

func add_sub() {

	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i: %v\n", i)

	atomic.AddInt32(&i, -1)
	fmt.Printf("i: %v\n", i)

}

func main() {
	var i int32 = 100
	atomic.LoadInt32(&i)
	fmt.Printf("ri: %v\n", ri)
}
