package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func test2() {

	tom := Person{
		101,
		"Tom",
		20,
	}

	pperson := &tom
	// fmt.Printf("tom: %v\n", tom)
	// fmt.Printf("pperson: %v\n", pperson)

	pperson.name = "SanYu"

	fmt.Printf("pperson: %p\n", pperson)

}

func main() {
	test2()
	person := Person{}
	fmt.Printf("person: %v\n", person)
}
