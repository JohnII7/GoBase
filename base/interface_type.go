package main

import "fmt"

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile struct {
}

// 单类多接口
func (m Mobile) playMusic() {
	fmt.Println("play muscic")
}

func (m Mobile) playVideo() {
	fmt.Println("play video")
}

// 多类单接口
type Pet interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}
func (cat Cat) eat() {
	fmt.Println("cat eat...")

}

func main() {
	m := Mobile{}
	m.playMusic()
	m.playVideo()

	var p Pet
	p = Dog{}
	p.eat()

	p = Cat{}
	p.eat()

}
