package main

import (
	"fmt"
)

func PopCount(x uint64) int {
	n := 0
	for i := uint64(0); i < 64; i++{
		if x & (1 << i) != 0 {
			n++
		}
	}
	return n
}

func main() {
	fmt.Printf("popcount: %v\n", PopCount(7))
}
