package main

import "fmt"

func main() {
	fmt.Println(f(5))

	fmt.Println(step(4))

}

func f(n int) int {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

//走台阶， 一次走一步，或者走两步
func step(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return step(n-1) + step(n-2)
}
