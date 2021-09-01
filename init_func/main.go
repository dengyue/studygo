package main

import "fmt"

var x = 10

const PI = 3.14

func init() {
	fmt.Println("execute automation")
	fmt.Println(x, PI)
}

func main() {
	fmt.Println("main function")
}
