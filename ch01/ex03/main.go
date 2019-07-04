package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%ds elapsed\n", time.Since(start).Nanoseconds())

	start = time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%d s elapsed\n", time.Since(start).Nanoseconds())


	start = time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Printf("%d s elapsed\n", time.Since(start).Nanoseconds())
}
