package main

import (
	"fmt"
	"reflect"
)

type Animal struct {

}

func (a *Animal) Eat(){
	fmt.Println("eat")
}

func main() {
	animal := Animal{}
	value := reflect.ValueOf(&animal)
	f := value.MethodByName("Eat")
	f.Call([]reflect.Value{})
}
