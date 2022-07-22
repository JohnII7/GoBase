package main

import "fmt"

type Person struct {
	name string
}

type Customer struct {
	name string
}

func (per Person) eat() {
	fmt.Printf("%v eat\n", per.name)
}

func (per Person) sleep() {
	fmt.Printf("%v sleep\n", per.name)
}

func (cunstomer Customer) login(name string, pwd string) bool {
	fmt.Printf("cunstomer.name: %v\n", cunstomer.name)
	if name == "John" && pwd == "123" {
		return true
	}
	return false
}

func main() {
	/* per := Person{
		name: "John",
	}
	per.eat()
	per.sleep() */

	cus := Customer{
		name: "John1",
	}
	f := cus.login(cus.name, "123")
	if f {
		fmt.Println("login!")
	} else {
		fmt.Println("login error")
	}
}
