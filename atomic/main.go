package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()

	// x++

	// lock.Lock()
	// x++
	// lock.Unlock()

	atomic.AddInt64(&x, 1)
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)

	x = 200
	ok := atomic.CompareAndSwapInt64(&x, 200, 100)
	fmt.Println(ok, x)

	ok := atomic.CompareAndSwapInt64(&, 200, 100)
	fmt.Println(ok, x)
}
