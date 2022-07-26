package main

import (
	"fmt"
	"os"
)

func openClose() {
	/* f, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("文件不存在")
	}
	fmt.Printf("f.Name(): %v\n", f.Name())
	f.Close() */

	// 如果文件不存在会创建
	f, err := os.OpenFile("a1.txt", os.O_RDONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}
}

func main() {
	openClose()
}
