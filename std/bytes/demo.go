package main

import "fmt"

func main() {
	var s string = "hello world"
	b1 := []byte{1, 2, 3, 4}
	s = string(b1)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("b1: %T\n", b1)
}
