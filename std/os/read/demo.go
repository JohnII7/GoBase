package main

import (
	"fmt"
	"io"
	"os"
)

func OpenClose() {
	// 打开文件，不能打开未cun
	/* f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("文件不存在: %v\n", err)
	}
	fmt.Printf("f.Name(): %v\n", f.Name()) */

	f, err := os.OpenFile("test.txt", os.O_RDONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f.Name())
		f.Close()
	}
}

// 读
func readOps() {
	f, _ := os.Open("test.txt")
	buf := make([]byte, 10)

	// f.Seek(3, 0)
	// 等同于 f.ReadAt(buf,3)，表示从3个字节开始读
	// 下一次读的位置, 第二个参数0表示相对文件开头，1表示相对文件当前位置，2表示相对文件末尾
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("buf: %v\n", buf)
		fmt.Printf("string(buf): %v\n", string(buf))
		fmt.Printf("n: %v\n", n)
	}
}

// 读目录
func readDir() {
	de, _ := os.ReadDir("..")
	for _, v := range de {
		fmt.Printf("v.Name(): %v\n", v.Name())
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
	}
}

func main() {
	OpenClose()
	// readOps()
	readDir()
}
