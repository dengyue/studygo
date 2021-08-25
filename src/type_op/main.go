package main

import "fmt"

type Person struct {
	name, gender string
}

func f(x Person) {
	x.gender = "女" // update the copy object
}

func f1(x *Person) {
	x.gender = "女" //update the original object
	//(*x).gender = "女"
}

//new -> return pointer, apply new memory for base type
//make-> return type,  apply new memory for slice, map and chan
func main() {
	var p Person
	p.name = "dengyue"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)

	f1(&p)
	fmt.Println(p.gender)

	var p2 = new(Person)
	p2.gender = "保密"
	fmt.Printf("%T \n", p2)
	fmt.Printf("%v  %x \n", p2, p2)

	var p3 = Person{
		name:   "dengyue",
		gender: "unknow",
	}
	fmt.Printf("%#v \n ", p3)

	p4 := Person{
		"zdy", "unkown",
	}
	fmt.Printf("%#v \n ", p4)
}
