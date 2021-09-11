package main

import (
	"fmt"
	"sync"
	"time"
)

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("B:%d\n", i)
	}
}

var wg sync.WaitGroup

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()

}
