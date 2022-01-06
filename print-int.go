package main

import (
	"os"
	"strconv"
)

func main() {
	print(42)
}

func print(i int) {
	s := strconv.FormatInt(int64(i), 10)
	os.Stdout.Write([]byte(s))
}
