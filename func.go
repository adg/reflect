package main

import (
	"fmt"
	"reflect"
)

// makeInc returns a function that takes the a value of the given
// type and returns that value incremented by one.
func makeInc(typ interface{}) interface{} {
	t := reflect.TypeOf(typ)
	ft := reflect.FuncOf([]reflect.Type{t}, []reflect.Type{t}, false)
	return reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		v := reflect.New(t).Elem()
		switch t.Kind() {
		case reflect.Int, reflect.Int64:
			v.SetInt(args[0].Int() + 1)
		case reflect.Float32, reflect.Float64:
			v.SetFloat(args[0].Float() + 1)
		}
		return []reflect.Value{v}
	}).Interface()
}

func main() {
	inc := makeInc(int(0)).(func(int) int)
	fmt.Println(inc(41))
}
