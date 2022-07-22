package main

import "fmt"

type Person struct {
	id   int
	name string
}

type Customer struct {
	id, age     int
	name, email string
}

func main() {
	tom := Person{
		id:   12,
		name: "ss",
	}
	fmt.Printf("tom: %v\n", tom)

	customer := Customer{
		id:    12,
		age:   22,
		name:  "cunstomer",
		email: "g@gmail.com",
	}
	fmt.Printf("customer: %v\n", customer)
}
