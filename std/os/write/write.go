package main

import (
	_ "fmt"
	"os"
)

func write() {
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0777)
	f.Write([]byte("GoGo"))
	f.Close()
}

func writeString() {
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0777)
	f.WriteString("this is a string")
	f.Close()
}

func writeAt() {
	f, _ := os.OpenFile("test.txt", os.O_RDWR, 0777)
	f.WriteAt([]byte(", "), 12)
	f.Close()
}
func main() {
	// write()
	// writeString()
	writeAt()
}
