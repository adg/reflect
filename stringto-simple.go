package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func stringTo(p interface{}, s string) {
	pv := reflect.ValueOf(p)
	if pv.Kind() != reflect.Ptr {
		panic("p must be a pointer")
	}
	switch e := pv.Elem(); e.Kind() {
	case reflect.Int:
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			e.SetInt(i)
		}
	case reflect.String:
		e.SetString(s)
	}
}

func main() {
	// start OMIT
	var i int
	stringTo(&i, "42")
	var s string
	stringTo(&s, "hello")
	fmt.Println(i, s)
	// end OMIT
}
