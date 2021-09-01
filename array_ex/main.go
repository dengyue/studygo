package main

import "fmt"

func main() {

	a1 := [...]int{1,2,3,4,5,6,7,8}
	sum := 0
	for _, v := range a1 {
		sum = sum + v
	}
	fmt.Println(sum)


	
}
