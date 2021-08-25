package main

import "fmt"

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	sum1(1, 2)
	sum2()
	sum3()

	f5()
	f6(1, 2, 3)
	f7("cd", 1, 2, 3, 4, 5, 6)
}

func sum(x int, y int) (s int) {
	s = x + y
	return
}

func sum1(x int, y int) {
	fmt.Println(x, y)
}

func sum2() {

}

func sum3() int {
	return 3
}

func f5() (int, string) {
	return 1, "chengdu"
}

func f6(x, y, z int) {
	fmt.Println(x, y, z)
}

func f7(s string, y ...int) {
	fmt.Println(s)
	fmt.Println(y)
}
