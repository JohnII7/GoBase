package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4}
	for i, v := range a1 {
		fmt.Printf("a1[%v]:%v\n", i, v)

	}
}
