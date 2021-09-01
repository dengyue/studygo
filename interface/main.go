package main

import (
	"fmt"
)

type car interface {
	driver()
}

type bus struct{}

type plane struct{}

func (b bus) driver() {
	fmt.Println("bus is slowly")
}

func (p plane) driver() {
	fmt.Println("plane is quickly")
}

func move(v car) {
	v.driver()
}

func main() {
	var bus = bus{}
	var plane = plane{}
	move(bus)
	move(plane)
}
