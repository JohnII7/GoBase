package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("age不能小于0")
	}
	return &Person{name: name, age: age}, nil
}

func main() {
	person, err := NewPerson("John", 20)
	if err == nil {
		fmt.Printf("person: %v\n", person)
	}
	if person.age > 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}
}
