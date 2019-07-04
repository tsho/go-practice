package main

import (
	"fmt"
)

func PopCount(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}

func main() {
	fmt.Printf("popcount: %v\n", PopCount(7))
}
