package main

var x = 100

func f1() int {
	x := 10
	println(x)
	return x
}

func main() {
	f1()
	f2(f1)
}

func f2(x func()int){
	ret := x()
	println(ret)
}
