package main

import (
	"bufio"
	"fmt"
	"strings"
)

func scanner() {
	s := strings.NewReader("Hello 世界")
	bs := bufio.NewScanner(s)
	//bs.Split(bufio.ScanRunes) 不会乱码
	bs.Split(bufio.ScanBytes)
	for bs.Scan() {
		fmt.Printf("bs.Text(): %v\n", bs.Text())
	}
}

func main() {
	scanner()
}
