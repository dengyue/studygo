package main

import "fmt"

func main() {
	d1 := newDog("dy")
	d1.wang()
}

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{name: name}
}

func (d dog) wang() {
	fmt.Printf("%s: wang, wang, wang", d.name)
}
