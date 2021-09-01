package main

import (
	"fmt"
)

func main() {

	f1 := func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("Hello World.")
	}
	f1(100, 200)

	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("Hello World.")
	}(1000, 2000)
}
