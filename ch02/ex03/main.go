package main

import (
	"fmt"

	"./popcount"
)

func main() {
	fmt.Printf("popcount: %v\n", popcount.PopCount(7))
	fmt.Printf("popcount loop: %v\n", popcount.PopCountL(7))
}
