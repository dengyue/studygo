package main

import "fmt"

type myInt int

type yourInt = int

func main() {
	var n myInt = 100
	fmt.Println(n)
	fmt.Printf("%T \n", n)

	var m yourInt = 100
	fmt.Println(m)
	fmt.Printf("%T \n", m)

}
