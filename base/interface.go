package main

import "fmt"

type USB interface {
	read()
	write()
}

type Mobile struct {
	model string
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Println("", c.name, "is reading")
}

func (c Computer) write() {
	fmt.Println("", c.name, "is writing")
}

func (m Mobile) read() {
	fmt.Println("", m.model, "is reading...")
}

func (m Mobile) write() {
	fmt.Println("", m.model, "is writing...")

}

func main() {
	c := Computer{
		name: "MacBook Pro",
	}
	c.read()
	c.write()

	m := Mobile{
		model: "5G",
	}
	m.read()
	m.write()
}
