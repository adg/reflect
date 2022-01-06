package main

import "os"

func main() {
	print("hello")
}

func print(s string) {
	os.Stdout.Write([]byte(s))
}
