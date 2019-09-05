package main

import (
	"fmt"

	. "./intset"
)

func main() {
	var x, y IntSet	
	x.Add(1)
	x.Add(2)
	x.Add(3)

	fmt.Println(x)
	fmt.Println(x.String())       // "{1 2 3}"
	fmt.Println(x.Len()) // "3"

	fmt.Println(x)
	x.Remove(2)
	fmt.Println(x)
	fmt.Println(x.String())

	y = * x.Copy()
	x.Clear()
	fmt.Println(x)  // "{}"
	fmt.Println(y) // "{1 3 42}"
}
