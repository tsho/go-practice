package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(countDiffHash(c1, c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

func countDiffHash(c1, c2 [32]byte) int {
	cnt := 0
	for i, _ := range c1 {
		cnt = countDiffBit(c1[i], c2[i])
	}
	return cnt
}

func countDiffBit(b1, b2 byte) int {
	cnt := 0
	for i := uint64(0); i < 8; i++ {
		bit1 := (b1 >> i) & 1
		bit2 := (b2 >> i) & 1
		if bit1 != bit2 {
			cnt++
		}
	}
	return cnt
}