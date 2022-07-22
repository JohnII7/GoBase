package main

import "fmt"

var i = initVar()

func init() {
	fmt.Println("init...")
}
func initVar() int {
	fmt.Println("initVar...")
	return 100
}

func main() {
	fmt.Printf("i: %v\n", i)
	fmt.Println("main...")
}
