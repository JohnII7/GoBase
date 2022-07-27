package main

import "fmt"

func appendSlice() {
	// 切片添加数字
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("i: %v\n", i)

	// 切片添加切片
	s1 := []int{11, 22, 33, 44}
	i2 := append(s, s1...)
	fmt.Printf("i2: %v\n", i2)
}

func lenSlice() {
	s := "hello world"
	fmt.Printf("len(s): %v\n", len(s))
}

func main() {
	appendSlice()
	lenSlice()

}
