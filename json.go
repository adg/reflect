package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// start OMIT
	var ints []int
	json.Unmarshal([]byte(`[1, 2, 3]`), &ints)
	var strings []string
	json.Unmarshal([]byte(`["one", "two", "three"]`), &strings)
	fmt.Println(ints, strings)
	// end OMIT
}
