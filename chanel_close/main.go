package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	x1, ok := <-ch
	fmt.Println(x1, ok)
	x2, ok := <-ch
	fmt.Println(x2, ok)
	x3, ok := <-ch
	fmt.Println(x3, ok)
}
