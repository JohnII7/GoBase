package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func marsh() {
	person := Person{
		Name:  "郭志宏的妈",
		Age:   20,
		Email: "GuosMom@gmail",
	}

	b, _ := xml.MarshalIndent(person, "", "")
	fmt.Printf("string(b): %v\n", string(b))
}

func main() {
	marsh()
}
