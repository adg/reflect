package main

import (
	"os"
	"reflect"
	"strconv"
)

// start OMIT

func sprint(in interface{}) string {
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10)
	default:
		return "???"
	}
}

func print(v interface{}) {
	os.Stdout.Write([]byte(sprint(v)))
}

// end OMIT

func main() {
	print(13)
	print("\n")
	type I int
	print(I(27))
}
