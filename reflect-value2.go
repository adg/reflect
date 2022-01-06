package main

import (
	"fmt"
	"reflect"
)

type S struct {
	A string
	B int
}

func main() {
	s := S{
		A: "hello",
		B: 42,
	}
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		ft := v.Type().Field(i)
		fv := v.Field(i)
		fmt.Println(ft.Name, fv.Interface())
	}
}
