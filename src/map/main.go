package main

import "fmt"

func main() {
	var m1 map[string]string
	fmt.Println(m1)
	m1 = make(map[string]string, 10)

	m1["sz"] = "shenzhen"
	m1["cd"] = "chengdu"
	fmt.Println(m1)

	fmt.Println(m1["cd"])

	v, ok := m1["test"]
	if !ok {
		fmt.Println("no key")
	} else {
		fmt.Println(v)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	for k := range m1 {
		fmt.Println(k)
	}

	for _, v := range m1 {
		fmt.Println(v)
	}

	delete(m1, "cd")
	fmt.Println(m1)
	delete(m1, "test")
	fmt.Println(m1)
}
