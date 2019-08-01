package main

import (
	"fmt"
)

const ( KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, KB*KB, KB*MB, KB*GB, KB*TB, KB*PB, KB*EB, KB*TB )

func main() {
	fmt.Printf("%b\n", KB)
	fmt.Printf("%b\n", MB)
	fmt.Printf("%b\n", GB)
	fmt.Printf("%b\n", TB)
	fmt.Printf("%b\n", PB)
	fmt.Printf("%b\n", EB)
	fmt.Printf("%b\n", ZB)
	fmt.Printf("%b\n", YB)
}
