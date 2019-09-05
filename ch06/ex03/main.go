package main

import (
	"fmt"

	. "./intset"
)

func main() {
	var x, y IntSet
	x.AddAll(1,4,7,100)
	y.AddAll(0,4,10,100,200)
	fmt.Println(x)
	fmt.Println(x.String())
	fmt.Println(y)
	fmt.Println(y.String())

	x.IntersectWith(&y)
	fmt.Println(x)
	fmt.Println(x.String())

	x.AddAll(1,4,7,100)
	x.DifferenceWith(&y)
	fmt.Println(x.String())

	x.AddAll(1,4,7,100)
	x.SymmetricDifference(&y)
	fmt.Println(x.String())

}