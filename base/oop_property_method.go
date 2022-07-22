package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (per Person) eat() {
	fmt.Println("eat...")
}

func (per Person) sleep() {
	fmt.Println("sleep...")
}

func (per Person) work() {
	fmt.Println("work...")
}

func main() {
	per := Person{
		name: "John",
		age:  24,
	}

	fmt.Printf("per: %v\n", per)
	per.eat()
	per.sleep()
	per.work()
}
