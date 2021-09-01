package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 100
	fmt.Printf("%T \n", n)
	fmt.Printf("%v \n", n)
	fmt.Printf("%d \n", n)
	fmt.Printf("%b \n", n)
	fmt.Printf("%o \n", n)
	fmt.Printf("%x \n", n)

	s1 := `
		离离原上
		一岁一枯荣
		`
	fmt.Println(s1)

	s2 := "理想" //一个中文占3个Bit
	s3 := "lixiang"

	ss := s2 + s3
	fmt.Println(ss)
	s2s3 := fmt.Sprintf("%s%s", s2, s3)
	fmt.Println(s2s3)

	fmt.Println(len(s2))
	fmt.Println(len(s3))

	fmt.Println(strings.Contains(s2s3, "理性"))
	fmt.Println(strings.Contains(s2s3, "理想"))

	fmt.Println(strings.HasPrefix(s2s3, "理想"))
	fmt.Println(strings.HasSuffix(s2s3, "理想"))

	s4 := "abcde"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "d"))

	s5 := "dengyue登月"

	for _, c := range s5 {
		fmt.Printf("%c\n ", c)
	}

	s6 := "白若不"
	s7 := []rune(s6)
	s7[0] = '红'
	fmt.Println(string(s7))

	c1 := "红"
	c2 := '红' // rune(int32)
	fmt.Printf("c1:%T, c2:%T \n ", c1, c2)

	//home work
	text := "hello我爱我的国"
	bt := []rune(text)
	count := 0
	for i := 0; i < len(bt); i++ {
		fmt.Println("%v", bt[i])
		if bt[i] > 127 {
			count++
		}
	}
	fmt.Printf("汉字个数: %d \n", count)

}
