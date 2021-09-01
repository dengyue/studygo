package main

import "fmt"

func main() {

	var m map[string]interface{}
	m = make(map[string]interface{}, 16)

	m["name"] = "dengyue"
	m["age"] = 10000
	m["merried"] = true
	m["hobby"] = [...]string{"chi", "he", "tiao"}

	fmt.Println(m)
}
