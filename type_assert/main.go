package main

import "fmt"

func assert(a interface{}) {
	fmt.Printf("%T\n", a)

	str, ok := a.(string)
	if !ok {
		fmt.Println("guess fault")
	} else {
		fmt.Println("this is a string :  ", str)
	}
}

func assert2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("this is a string : ", t)
	case int:
		fmt.Println("this is int :  ", t)
	case bool:
		fmt.Println("this is a bool: ", t)
	}

}

func main() {
	// assert(100)
	assert2(100)
}
