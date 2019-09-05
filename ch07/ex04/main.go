package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF // Must set EOF, otherwise it does not end
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}


func main() {
	doc, _ := html.Parse(NewReader("<html><body><h1>hello</h1></body></html>"))
	fmt.Println(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
}