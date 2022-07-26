package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func createDir() {
	// 创建单级目录
	err := os.Mkdir("testDir", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// 创建多级目录
	err2 := os.MkdirAll("multiDir/a/b", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除目录或者文件
func remove() {
	// 只能删除单级目录
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	os.Remove("testDir")
}

func removeAll() {
	os.RemoveAll("multiDir")
}

// 获取当前目录
func wd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
}

// 修改工作目录
/* func chWD(){
	err := os.Chdir("d:/")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())
} */

// 获取临时目录
func getTemp() {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

// 重命名
func rename() {
	err := os.Rename("test.txt", "text2.txt")
	if err != nil {
		fmt.Printf("没有可用于重命名的文件: %v\n", err)
	}
}

// 读文件
func readFile() {
	b, _ := os.ReadFile("test3.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

// 写文件
func writeFile() {
	os.WriteFile("test4.txt", []byte("Hello,World and Go"), os.ModePerm)
}

func main() {
	createFile()
	createDir()
	// remove()
	removeAll()
	wd()
	getTemp()
	rename()
	readFile()
	writeFile()
}
