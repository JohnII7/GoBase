package main

import "fmt"

type Person struct {
	name string
}

func main() {
	p1 := Person{"tom"}

	p2 := &Person{"tom"}
	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p1: %T\n", p2)

	showPerson1(p1)
	showPerson2(p2)

	showPerson(p1)
	showPerson(*p2)
}
func showPerson1(per Person) {
	per.name = "值传递"
}
func showPerson2(per *Person) {
	per.name = "值传递"
}
func showPerson(per Person) {
	fmt.Printf("per.name: %v\n", per.name)
}
