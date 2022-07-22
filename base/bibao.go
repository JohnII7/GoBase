package main

import (
	"fmt"
	"strings"
)

func main() {
	// 闭包 = 函数 + 外层变量的引用
	r := a("沙河小王房子") // r 就是一个闭包
	r()              // 相当于执行了a函数内部的匿名函数

	f := makeSuffixFunc(".txt")
	res := f("沙河娜扎")
	fmt.Printf("res: %v\n", res)

	x, y := calc(100)
	resx := x(200)
	fmt.Printf("resx: %v\n", resx)

	resy := y(1)
	fmt.Printf("resy: %v\n", resy)

}

// 返回值是一个函数
func a(name string) func() {
	return func() {
		fmt.Println("hello ", name)
	}
}

// 举例
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(b int) int {
		base -= b
		return base
	}
	return add, sub
}
