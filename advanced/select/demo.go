package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanStr <- "Hello"
		// close(chanInt)
		// close(chanStr)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt: %v\n", r)
			fmt.Printf("chanInt r type: %T\n", r)

		case r := <-chanStr:
			fmt.Printf("chanStr: %v\n", r)
			fmt.Printf("chanStr r type: %T\n", r)

		default:
			fmt.Println("default: end")
		}

		time.Sleep(time.Second)
	}
}
