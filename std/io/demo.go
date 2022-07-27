package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func testCopy() {
	r := strings.NewReader("Hello World\n")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	testCopy()
}
