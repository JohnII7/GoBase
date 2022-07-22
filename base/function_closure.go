package main

import "fmt"

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	f := add()
	r := f(10)
	fmt.Printf("r: %v\n", r)

	r = f(20)
	fmt.Printf("r: %v\n", r)

	r = f(30)
	fmt.Printf("r: %v\n", r)

	fmt.Println("==========================")
	f1 := add()
	r = f1(100)
	fmt.Printf("r: %v\n", r)
}
