//https://golang.org/pkg/unicode/#IsSpace

package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "\t\ts \t\t\t\t \n\ntest"
	fmt.Println(s)
	fmt.Println(string(removeDuplicateSpace([]byte(s))))
}

func removeDuplicateSpace(s []byte) []byte {
	var buf bytes.Buffer
	lastUnicodeSpace := false
	for i := 0; i < len(s) - 1; {
		r, size := utf8.DecodeRuneInString(string(s[i:]))
		if unicode.IsSpace(r) && lastUnicodeSpace {
			i+=size
			continue
		} else if unicode.IsSpace(r) && !lastUnicodeSpace {
			buf.WriteRune(' ')
			lastUnicodeSpace = true
		} else {
			buf.WriteRune(r)
			lastUnicodeSpace = false
		}
		i+=size
	}
	return buf.Bytes()
}