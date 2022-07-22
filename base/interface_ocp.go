package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}
type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("DOG EAT")
}

func (dog Dog) sleep() {
	fmt.Println("DOG SLEEP")
}

func (cat Cat) eat() {
	fmt.Println("CAT EAT")
}

func (cat Cat) sleep() {
	fmt.Println("CAT SLEEP")
}

type Person struct {
}

func (person Person) care(pet Pet) {
	pet.eat()

	pet.sleep()
}

func main() {
	dog := Dog{}
	cat := Cat{}
	person := Person{}
	person.care(cat)
	person.care(dog)
}
