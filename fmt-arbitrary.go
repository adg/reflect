package main

import "fmt"

func main() {
	// start OMIT
	i := 42
	s := "hello"
	f := 3.14159
	type Vector struct{ X, Y int }
	v := Vector{3, 4}
	m := map[Vector][]int{
		{1, 2}: {3, 4},
		{5, 6}: {7, 8},
	}
	fmt.Println(s, i, f, v, m)
	// end OMIT
}
