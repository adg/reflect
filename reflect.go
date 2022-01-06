package main

import (
	"os"
	"reflect"
	"strconv"
)

func main() {
	// start OMIT
	var i interface{} = 42
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	print("(")
	print(t.String())
	print(", ")
	print(v.Int())
	print(")")
	// end OMIT
}

func print(v interface{}) {
	var out string
	switch v := v.(type) {
	case string:
		out = v
	case int64:
		out = strconv.FormatInt(int64(v), 10)
	default:
		out = "???"
	}
	os.Stdout.Write([]byte(out))
}
