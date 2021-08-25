package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var dengyue Person

	dengyue.name = "dengyue"
	dengyue.age = 100
	dengyue.gender = "男"
	dengyue.hobby = []string{"basketball", "football"}

	fmt.Println(dengyue)

	var s struct {
		name string
		age  int
	}
	s.name = "yusheng"
	s.age = 9

	fmt.Printf("%T %v", s, s)
}
