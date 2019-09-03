package main

import (
	"fmt"
)

func noreturn(i string) (o string) {
	o = i
	defer func() { recover() }()
	panic(i)
}

func main() {
	fmt.Println(noreturn("Hello, World"))
}