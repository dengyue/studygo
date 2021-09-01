package main

import "fmt"

func main() {

	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1: %T, a2： %T \n ", a1, a2)

	a1 = [3]bool{true,false,false}
	fmt.Println(a1)

	a10 := [...]int {0,1,2,3,4,5,6,7,8,9}
	fmt.Println(a10)

	a3 := [5]int {0:1, 4:2}
	fmt.Println(a3)

	cities := [3]string{"北京","成都","上海"}
	for i := 0; i< len(cities);i++{
		fmt.Println(cities[i])
	}

	for i, v := range cities{
		fmt.Println(i,v)
	}

	//数组是值类型
	b1 := [3]int{1,2,3}
	b2 := b1
	b2[0] =  100
	fmt.Println(b1, b2)

}
