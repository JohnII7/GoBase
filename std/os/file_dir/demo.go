package main

import (
	"fmt"
	"os"
)

// 创建文件

func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

func main() {
	createFile()
}
