package main

import (
	"fmt"
	"sync"
)

var b chan int

var wg sync.WaitGroup

func main() {
	b = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c := <-b
		fmt.Println(c)
	}()
	b <- 10
	wg.Wait()
	fmt.Println("main")
}
