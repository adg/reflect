package main

import (
	"os"
	"strconv"
)

func print(v interface{}) {
	var out string
	if s, ok := v.(string); ok {
		out = s
	} else if i, ok := v.(int); ok {
		out = strconv.FormatInt(int64(i), 10)
	} else {
		out = "???"
	}
	os.Stdout.Write([]byte(out))
}

func main() {
	print("hello\n")
	print(42)
	type I int
	print(I(42))
}
