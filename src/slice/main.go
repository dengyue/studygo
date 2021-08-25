package main

import "fmt"

func main() {
	var s1 []string
	var s2 []string
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	s1 = []string{"abc", "def"}
	s2 = []string{"a", "b", "c"}
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("len(s1)=%d cap(s1)=%d \n", len(s1), cap(s1))

	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4]
	fmt.Println(s3)
	fmt.Printf("len(s3)=%d cap(s3)=%d \n", len(s3), cap(s3))

	s6 := a1[3:]
	fmt.Println(s6)
	fmt.Printf("len(s6)=%d cap(s6)=%d \n", len(s6), cap(s6))

	s7 := a1[:]
	fmt.Println(s7)
	fmt.Printf("len(s7)=%d cap(s7)=%d \n", len(s7), cap(s7))

	s7[0] = 100
	fmt.Println(s7)
	fmt.Println(a1)
	fmt.Println(s3)

	s7[3] = 1000
	fmt.Println(s7)
	fmt.Println(a1)
	fmt.Println(s3)
	fmt.Println(s6)

}
