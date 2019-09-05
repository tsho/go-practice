package main

import (
	"bytes"
	"fmt"
)

func VariadicJoin(sep string, a ...string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := bytes.Buffer{}
	for _, s := range a[:len(a)-1] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	b.WriteString(a[len(a)-1])
	return b.String()
}

func main() {
	fmt.Println(VariadicJoin("/", "test", ".", "com"))
	fmt.Println(VariadicJoin("/", "test", "."))
}