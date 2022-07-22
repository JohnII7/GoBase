package main

import "fmt"

func main() {
	a := 1
	f1(a)
	fmt.Printf("a: %v\n", a)

	s := []int{1}
	f2(s)
	fmt.Printf("s: %v\n", s)
}

func f1(a int) {
	a = 100
}

func f2(s []int) {
	s[0] = 100
}
