package main

func main() {
	println(f1())
	println(f2())
	println(f3())
	println(f4())
}


func f1() int{
	x := 5
	defer func(){
		x++
	}()
	return x
}

func f2() (x int){
	defer func(){
		x++
	}()
	return 5
}
func f3() (y int){
	x := 5
	defer func(){
		x++
	}()
	return x
}

func f4() (x int){
	defer func(x int) {
		x++
	}(x)
	return 5
}