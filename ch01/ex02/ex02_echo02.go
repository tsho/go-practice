// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", " "
	for i := 0; i < len(os.Args); i++ {
		s = strconv.Itoa(i) + sep + os.Args[i]
		fmt.Println(s)
	}
}

//!-
