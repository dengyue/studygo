package main

import "fmt"

func main() {
	addf := adder(100)
	sum := addf(200)
	fmt.Println(sum)
}

//闭包=一个匿名函数+一个外面变量的应用
func adder(x int) func(int) int {
	return func(i int) int {
		x += i
		return x
	}
}
