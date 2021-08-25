package main

import "fmt"

func main() {
	f1()
	f2("dengyue")
}

func f1() {
	fmt.Println("it's f1")
}

func f2(name string) {
	fmt.Println("Hello, ", name)
}
