package main

import (
	"os"
	"reflect"
	"strconv"
)

func print(v interface{}) {
	os.Stdout.Write([]byte(sprint(v)))
}

type S struct {
	A string
	B int
}

// start OMIT

func sprint(in interface{}) string {
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Struct:
		out := "{"
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				out += ", "
			}
			out += v.Type().Field(i).Name
			out += ": "
			out += sprint(v.Field(i).Interface())
		}
		out += "}"
		return out
	// (string, int, and default cases as before)
	case reflect.String:
		return v.String()
	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10)
	default:
		return "???"
	}
}

func main() {
	print(S{A: "hello", B: 42})
}
