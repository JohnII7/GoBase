package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3}
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))
	fmt.Printf("s1: %v\n", s1)
}
