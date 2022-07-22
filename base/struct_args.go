package main

import "fmt"

type Person struct {
	id   int
	name string
}

func main() {
	p := &Person{
		1,
		"John",
	}
	fmt.Printf("p: %v\n", p)
	fmt.Printf("addr_structFunc(p): %v\n", addr_structFunc(*p))

	fmt.Printf("p: %v\n", p)
	fmt.Printf("structFunc(p): %v\n", structFunc(p))
	fmt.Printf("p: %v\n", p)

}

func addr_structFunc(person Person) Person {
	person.id = -1
	person.name = "值传递"
	a := *&person
	return a
}

func structFunc(person *Person) *Person {
	person.id = 11
	person.name = "地址传递"
	a := *&person
	return a
}
