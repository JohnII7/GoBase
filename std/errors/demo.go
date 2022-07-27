package main

import (
	"errors"
	"fmt"
	_ "os"
)

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空\n")
		return "nil", err
	} else {
		return s, nil
	}
}

func main() {
	s, err := check("")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	} else {
		fmt.Printf("s: %v\n", s)
	}

	//=======================================

	s, err = check("Hello")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	} else {
		fmt.Printf("s: %v\n", s)
	}
}
