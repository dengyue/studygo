package main

import (
	"fmt"
	"strings"
)

func main() {

	var s1 = make([]map[int]string, 1, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][100] = "test"
	s1[0][101] = "test101"
	fmt.Println(s1)

	var m1 = make(map[string][]int, 10)
	m1["cd"] = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(m1)

	//split
	var w1 = "how do you do"
	cw := strings.Split(w1, " ")
	r1 := make(map[string]int, 10)
	for _, w := range cw {
		if _, ok := r1[w]; !ok {
			r1[w] = 1
		} else {
			r1[w]++
		}
	}
	fmt.Println(r1)

	//上海自来水来自海上
	var ss = "上海自来水来自海上"
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("no 回文")
			return
		}
	}
}
