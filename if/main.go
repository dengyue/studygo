package main

import "fmt"

func main() {

	if age := 20; age > 35 {
		fmt.Println("中年")
	} else if age > 20 {
		fmt.Println("青年")
	} else {
		fmt.Println("少年")
	}
}
