package main

import (
	"os"
	"strconv"
)

func print(v interface{}) {
	var out string
	switch v := v.(type) {
	case string:
		out = v
	case int:
		out = strconv.FormatInt(int64(v), 10)
	default:
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
