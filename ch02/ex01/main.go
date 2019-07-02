package main

import (
	"fmt"
	"./tempconv"
)

func main() {
	fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroK)
	fmt.Printf("0K %v\n", tempconv.CToK(0))
}