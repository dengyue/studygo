package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var scoremap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoremap[key] = value
	}
	fmt.Println(scoremap)

	var keys = make([]string, 0, 200)
	for key := range scoremap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoremap[key])
	}
}
