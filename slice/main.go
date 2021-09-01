package main

import "fmt"

func main() {
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//initial
	s1 = []int{1, 2, 3}
	s2 = []string{"beijing", "shanghai", "chengdu"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//length and capacity
	fmt.Printf("len(s1): %d, cap(s1): %d \n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d, cap(s2): %d \n", len(s2), cap(s2))

	//数组切片
	a1 := [...]int{1,3,5,7,9,11,13}
	s3 := a1[0:4] //左闭右开
	fmt.Println(s3)
	//切片的容量是底层数组的容量
	fmt.Printf("len(s3): %d, cap(s3): %d \n", len(s3), cap(s3))


	//切片是传地址（引用）
	s3[0]= 15
	fmt.Println(s3)
	fmt.Println(a1)

}
