package main

import (
	"fmt"
	"reflect"
)

func main() {
	// mark1 OMIT
	v1 := reflect.ValueOf(13)
	fmt.Println(v1.Type().Name(), v1.Kind())

	type I int
	v2 := reflect.ValueOf(I(27))
	fmt.Println(v2.Type().Name(), v2.Kind())
	// mark2 OMIT
	fmt.Println(v1.Int() + v2.Int())
	// mark3 OMIT
	i1 := v1.Interface().(int)
	i2 := v2.Interface().(I)
	fmt.Println(i1, i2)
	// end OMIT
}
