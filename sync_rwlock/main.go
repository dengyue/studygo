package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	defer wg.Done()

	// lock.Lock()
	rwLock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwLock.Unlock()
}

func write() {
	defer wg.Done()

	// lock.Lock()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Since((start)))
}
