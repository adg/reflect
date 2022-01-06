package main

import "fmt"

type Doubler int

func (d Doubler) Compute() int { return int(d) * 2 }

func main() {
	var f func(Doubler) int = Doubler.Compute
	fmt.Print(f(2))
}
