package main

import "fmt"

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}

// add
func test1() {
	s1 := []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v\n", s1)
}

// del
func test2() {
	s1 := []int{1, 2, 3, 4, 5}
	// 删除 i 2
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
}

// update
func test3() {
	s1 := []int{1, 2, 3, 4}
	s1[1] = 100
	fmt.Printf("s1: %v\n", s1)
}

// query
func test4() {
	s1 := []int{1, 2, 3, 4}
	key := 2
	for i, v := range s1 {
		if v == key {
			fmt.Printf("i: %v\n", i)
		}
	}
}
