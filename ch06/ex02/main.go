package main

import (
	"fmt"

	. "./intset"
)

func main() {
	var x IntSet	
	x.AddAll(1, 2, 3)
	fmt.Println(x.String())       // "{1 2 3}"
}
