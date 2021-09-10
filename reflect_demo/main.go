package main

import (
	"fmt"
	"math"
	"reflect"
)

type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
	fmt.Printf("type name: %v kind:%v \n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))

	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a float32 = math.Pi
	reflectType(a)

	var b int64 = 100
	reflectType(b)

	var c = Cat{}
	reflectType(c)

	//value
	reflectValue(a)
	reflectValue(b)

	//set value
	reflectSetValue(&b)
	fmt.Println(b)
}
