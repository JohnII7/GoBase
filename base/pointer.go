package main

import "fmt"

func main() {
	var ip *int
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %T\n", ip)

	i := 100
	ip = &i
	add := &i

	fmt.Printf("ip: %v\n", ip)
	fmt.Println("ip:", *ip)
	fmt.Println("add", *add)
}
