package main

import "fmt"

func main() {
	// test1()
	test2()
}

func test1() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[0:3]
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:]
	fmt.Printf("s3: %v\n", s3)
	s4 := s1[2:5]
	fmt.Printf("s4: %v\n", s4)
}

func test2() {
	var a1 = [...]int{1, 2, 3, 4, 5, 6}
	a2 := a1[:3]
	fmt.Printf("a2: %v\n", a2)
}
