package main

import "fmt"

func main() {
	test1()
	test2()
}

// 浅拷贝指拷贝地址, 会改变原来切片的值

func test1() {
	fmt.Println("test1===========")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1[1] = 1000
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

// copy不会更改原来数据, 即深拷贝
func test2() {
	fmt.Println("test2===========")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5)
	copy(s2, s1)
	s2[1] = 100000
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

}
