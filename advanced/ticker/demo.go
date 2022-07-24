package main

import (
	"fmt"
	"time"
)

func main() {
	chanInt := make(chan int)
	ticker := time.NewTicker(time.Second)

	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:

			case chanInt <- 2:

			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Println("收到:", v)
		sum += v
		if sum >= 10 {
			break
		}
	}
}
