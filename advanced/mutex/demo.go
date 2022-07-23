package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wg sync.WaitGroup

var lock sync.Mutex

func add() {
	defer wg.Done()
	i += 1
	fmt.Printf("i++: %v\n", i)
	time.Sleep(time.Second)
}

func sub() {
	defer wg.Done()
	time.Sleep(time.Second)
	i -= 1
	fmt.Printf("i--: %v\n", i)
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}

	wg.Wait()

	fmt.Println("end i:", i)
}
