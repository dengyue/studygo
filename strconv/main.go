package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "1000"

	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("error, error: %v\n", err)
		return
	}
	fmt.Println(i)

	m, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(m)

	n := 9000
	fmt.Printf("value %s, type is %T \n", strconv.Itoa(n), strconv.Itoa(n))
}
