package main

import "fmt"

var c = make(chan int)

func main() {
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		// 如果不加close关闭, 会发生死锁
		close(c)
	}()

	for i := 0; i < 3; i++ {
		r := <-c
		fmt.Printf("r: %v\n", r)
	}

}
