package main

import (
	"fmt"
)

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat...")
}

func (a Animal) sleep() {
	fmt.Println("sleep...")
}

type Dog struct {
	Animal
	color string
}

type Cat struct {
	Animal
	action string
}

func main() {
	dog := Dog{
		Animal{name: "dog", age: 2},
		"yellow",
	}

	dog.eat()
	dog.sleep()
	fmt.Printf("dog.color: %v\n", dog.color)

	cat := Cat{
		Animal{name: "dog", age: 2},
		"后空翻",
	}

	cat.eat()
	cat.sleep()
	fmt.Printf("cat.action: %v\n", cat.action)
}
