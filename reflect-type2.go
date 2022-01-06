package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(S{})
	fmt.Println(t.Name(), t.Kind())
	fmt.Println("fields:", t.NumField(), "methods:", t.NumMethod())
	f0, f1 := t.Field(0), t.Field(1)
	fmt.Println("fields:", f0.Name, f0.Type, f1.Name, f1.Type)
	m := t.Method(0)
	fmt.Println("method:", m.Name, "args:", m.Type.NumIn(), "returns:", m.Type.NumOut())
	fmt.Println("args:", m.Type.In(0), m.Type.In(1))
	fmt.Println("returns:", m.Type.Out(0))
}

type S struct {
	A string
	B int
}

func (S) Foo(string) int { return 0 }
