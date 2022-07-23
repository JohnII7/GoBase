package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func showMsg(i int) {

	// 等于wg.Add(-1)
	defer wg.Done()
	fmt.Printf("i: %v\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go showMsg(i)
		wg.Add(1)
	}

	// 主协程
	wg.Wait()
	fmt.Println("end")
}
