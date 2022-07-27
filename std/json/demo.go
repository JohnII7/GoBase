package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func marsh() {
	p := Person{
		Name:  "郭宏志的妈",
		Age:   20,
		Email: "guodema@gmail.com",
	}

	b, _ := json.Marshal(p)
	fmt.Printf("string(b): %v\n", string(b))
}
func unmarsh() {
	b := []byte(`{"Name":"郭宏志的妈","Age":20,"Email":"guodema@gmail.com"}`)
	var p Person
	json.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

func main() {
	marsh()
	unmarsh()
}
