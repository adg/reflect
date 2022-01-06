package main

import "fmt"

type Computer interface {
	Compute() int
}

type Doubler int

func (d Doubler) Compute() int { return int(d) * 2 }

type Halver int

func (h Halver) Compute() int { return int(h) / 2 }

func main() {
	var c Computer = Doubler(2)
	fmt.Print(c.Compute())
}
