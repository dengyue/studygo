package main

import "fmt"

func main() {
	s1 := []string{"beijing", "shanghai", "shenzhen"}
	// s1[2] = "guangzhou"
	// fmt.Println(s1)
	s1 = append(s1, "guangzhou", "chengdu")
	fmt.Println(s1)

	s2 := []string{"nanchong", "mianyang"}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
