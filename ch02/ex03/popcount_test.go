package main

import (
	"fmt"
	"testing"
	"./popcount"
)

func BenchmarkPop(b *testing.B){
	fmt.Printf("popcount: %v\n", popcount.PopCount(0))
	fmt.Printf("popcount: %v\n", popcount.PopCount(8))
	fmt.Printf("popcount: %v\n", popcount.PopCount(254))
	fmt.Printf("popcount: %v\n", popcount.PopCountL(2048))
}

func BenchmarkPopl(b *testing.B){
	fmt.Printf("popcount: %v\n", popcount.PopCountL(0))
	fmt.Printf("popcount: %v\n", popcount.PopCountL(8))
	fmt.Printf("popcount: %v\n", popcount.PopCountL(254))
	fmt.Printf("popcount: %v\n", popcount.PopCountL(2048))
}
