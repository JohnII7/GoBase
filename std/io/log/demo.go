package main

import (
	"fmt"
	"log"
)

func logPrint() {
	log.Print("Log\n")
	log.Printf("log format %d\n", 100)
	name := "John"
	age := 24
	log.Println(name, "", age)

}

func panicPrint() {
	// panic结束后 后面的代码不会在执行
	defer fmt.Println("panic结束之前的defer")
	log.Panic("panic exe")
}

func fatalPrint() {
	// fatal之后不会执行defer
	defer fmt.Println("log fatal defer")
	log.Fatal("log fatal")

}

func main() {
	logPrint()
	// panicPrint()
	fatalPrint()
}
