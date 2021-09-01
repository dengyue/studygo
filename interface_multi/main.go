package main

import "fmt"

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}
type eater interface {
	eat(string)
}

type Cat struct {
	name string
	feet uint8
}

func (c *Cat) move() {
	fmt.Println("zou mao bu.")
}

func (c *Cat) eat(food string) {
	fmt.Printf("mao chi %s \n", food)
}

func main() {
	var c = Cat{
		name: "mao",
		feet: 4,
	}
	var m mover
	m = &c
	fmt.Println(m)
}
