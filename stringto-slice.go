package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
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
	// start OMIT
	// stringTo function as before...
	case reflect.Slice:
		ss := strings.Fields(s)
		slice := reflect.MakeSlice(e.Type(), len(ss), len(ss))
		for i, s := range ss {
			stringTo(slice.Index(i).Addr().Interface(), s)
		}
		e.Set(slice)
		// end OMIT
	}
}

func main() {
	// start2 OMIT
	var i []int
	stringTo(&i, "1 1 2  3   5     8")
	var s []string
	stringTo(&s, "what a lovely    day")
	fmt.Println(i, s)
	// end OMIT
}
