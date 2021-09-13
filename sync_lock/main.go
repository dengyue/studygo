package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

func add() {
	defer wg.Done()
	for i := 0; i<5000000; i++{
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}