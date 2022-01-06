package main

import (
	"fmt"
	"reflect"
)

func main() {
	// start OMIT
	t1 := reflect.TypeOf("hello")
	fmt.Println(t1.Name(), t1.Kind())

	t2 := reflect.TypeOf(42)
	fmt.Println(t2.Name(), t2.Kind())

	type I int
	t3 := reflect.TypeOf(I(42))
	fmt.Println(t3.Name(), t3.Kind())

	fmt.Println(t1.Kind() == t2.Kind())
	fmt.Println(t2.Kind() == t3.Kind())
	// end OMIT
}
