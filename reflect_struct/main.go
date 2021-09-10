package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	p1 := person{
		Name: "xiaowangzi",
		Age:  9000,
	}
	t := reflect.TypeOf(p1)
	fmt.Println(t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	if ageField, ok := t.FieldByName("Age"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", ageField.Name, ageField.Index, ageField.Type, ageField.Tag.Get("json"))
	}

}
