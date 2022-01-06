package main

import "fmt"

func main() {
	var i interface{}
	// start OMIT
	i = "hello"
	fmt.Println(i)
	i = 42
	fmt.Println(i)
	// end OMIT
}
