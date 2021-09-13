package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

var m2 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	//if the i is 20, there isn't fatal error
	//if the i is greater than 20, then it will report error.
	// for i := 0; i < 21; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		key := strconv.Itoa(n)
	// 		set(key, n)
	// 		fmt.Printf("k=%v, v=%v\n", key, get(key))
	// 		defer wg.Done()
	// 	}(i)
	// }

	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("k=%v, v=%v\n", key, value)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
