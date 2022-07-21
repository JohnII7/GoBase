package main

import "fmt"

type WebSite struct {
	Name string
}

func main() {
	// 普通占位符
	// %v var 对应值的默认格式
	site := WebSite{Name: "duoke360"}
	fmt.Printf("site: %v\n", site)

	// %#v 对应值的Go语法表示
	fmt.Printf("site:%#v\n", site)

	// %T 对应值的类型Go语法表示
	fmt.Printf("site:%T\n", site)

	//=================================

	// 布尔占位符, 实际和%v效果一样
	b := true
	fmt.Printf("b: %t\n", b)
	fmt.Printf("b: %v\n", b)

	//=================================

	// 整形占位符

	i := 8
	fmt.Printf("i: %v\n", i)
	// 十进制
	fmt.Printf("i: %d\n", i)
	// 二进制
	fmt.Printf("i: %b\n", i)
	i = 65
	// 字符
	fmt.Printf("i: %c\n", i)
	// 十六进制 a-f
	fmt.Printf("i: %x\n", 12312)
	// 十六进制 A-F
	fmt.Printf("i: %X\n", i)
	// unicode
	fmt.Printf("i: %U\n", i)

}
