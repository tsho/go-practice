package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)

	fmt.Println(s)
}

func rotate (s []int, n int) {
	tmp := append(s, s...)
	copy(s, tmp[n:n+len(s)])
}
