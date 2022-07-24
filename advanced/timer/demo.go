package main

import (
	"fmt"
	"time"
)

func main() {
	// timer1 := time.NewTimer(time.Second * 3)
	// fmt.Printf("time.Now(): %v\n", time.Now())
	// t1 := <-timer1.C
	// fmt.Printf("t1: %v\n", t1)

	// fmt.Printf("time.Now(): %v\n", time.Now())
	// timer := time.NewTimer(time.Second * 2)
	// <-timer.C
	// fmt.Printf("time.Now(): %v\n", time.Now())

	timer := time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Println("func...")
	}()
	s := timer.Stop()
	if s {
		fmt.Println("stop")
	}

	time.Sleep(time.Second * 2)
	fmt.Println("main end")
}
