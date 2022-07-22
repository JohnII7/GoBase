package main

import "fmt"

// 函数可以作为参数
func sayHello(name string) {
	fmt.Println("Hello", name)
}

func test(name string, f func(string)) {
	f(name)
}

// 也可以作为返回值
func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
func cal(oprator string) func(int, int) int {
	switch oprator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	// test("tom", sayHello)
	ff := cal("+")
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)

	ff = cal("-")
	r = ff(2, 2)
	fmt.Printf("r: %v\n", r)
}
