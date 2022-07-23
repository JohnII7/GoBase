package main

import "fmt"
import "runtime"
import "time"

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("a: %v\n", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("b: %v\n", i)
	}
}

func main() {
	// 查看逻辑CPU(线程数)
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	go a()
	go b()

	time.Sleep(time.Second)
}
