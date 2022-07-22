package main

import "fmt"

func main() {
	sum := sum(1, 2)
	fmt.Printf("sum: %v\n", sum)

	max := comp(111, 22)
	fmt.Printf("max: %v\n", max)
}

func sum(a int, b int) int {
	return a + b
}

func comp(a int, b int) int {
	if a > b {
		return a
	}
	return b

}
