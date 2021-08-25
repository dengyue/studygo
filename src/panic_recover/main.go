package main

import "fmt"

func main() {
	funcA()
	funcB()
	funcC()
}

func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("relese resource")
	}()
	panic("error and need to release resource") //crashed, then exit
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}
