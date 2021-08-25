package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 4)
	copy(a3, a1)

	fmt.Println(a2)
	fmt.Println(a3)

	a1[0] = 100
	fmt.Println(a1, a2, a3)

	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)

	fmt.Println(cap(a1))

	x1 := [...]int{1, 2, 3}
	x11 := x1[:]
	fmt.Printf("%T, %T", x1, x11)
}
