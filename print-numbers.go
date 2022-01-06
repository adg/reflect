package main

import (
	"os"
	"reflect"
	"strconv"
)

func print(in interface{}) {
	var out string
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.String:
		out = v.String()
	case reflect.Int:
		out = strconv.FormatInt(v.Int(), 10)
	default:
		out = "???"
	}
	os.Stdout.Write([]byte(out))
}

func main() {
	print(13)
	print("\n")
	type I int
	print(I(27))
}
