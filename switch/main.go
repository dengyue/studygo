package main

import "fmt"

func main() {

	//reduce volume if/else
	var n = 5
	switch n {
	case 1:
		fmt.Println(n)
	case 2:
		fmt.Println(n)
	case 3:
		fmt.Println(n)
	case 4:
		fmt.Println(n)
	case 5:
		fmt.Println(n)
	default:
		fmt.Println("other")
	}

	switch n := 8; n {
	case 8:
		fmt.Println(8)
		fallthrough
	case 9:
		fmt.Println(9)
	default:
		fmt.Println("default")
	}
}
