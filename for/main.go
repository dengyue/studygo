package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for {
	//	fmt.Println("aaaaa")
	//	break
	//}

	//for range
	s := "hello 登月"
	for i, v := range s {
		fmt.Printf("%d, %c \n ", i, v)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	//home work
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d*%d=%d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
