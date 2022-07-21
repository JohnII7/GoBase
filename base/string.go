package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString("don't")
	buffer.WriteString("kill")

	fmt.Printf("buffer String : %s\n", buffer.String())

	// 字符串函数
	s := "Hello World"
	fmt.Printf("len(s): %v\n", s)

	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))

	fmt.Printf("strings.Contains(s, \"Hello\"): %v\n", strings.Contains(s, "Hello"))

	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))

	fmt.Printf("strings.ToUpper(s): %v\n", strings.ToUpper(s))

	fmt.Printf("strings.HasPrefix(s, \"H\"): %v\n", strings.HasPrefix(s, "H"))

	fmt.Printf("strings.HasSuffix(s, \"rld\"): %v\n", strings.HasSuffix(s, "rld"))

	fmt.Printf("strings.Index(s, \"ll\"): %v\n", strings.Index(s, "ll"))

	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))
}
