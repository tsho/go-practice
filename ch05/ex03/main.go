// Exercise 5.3
// Display text from html structure

package main

import (
	"fmt"
	"os"
	"unicode"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	visit(doc)
}

func visit(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode && n.Parent.Data != "script" {
		if skipSpace([]rune(n.Data)) {
			fmt.Println(n.Data)
		}
	}
	visit(n.FirstChild)
	visit(n.NextSibling)

}

func skipSpace(str []rune) bool {
	flg := false
	for _, s := range str {
		if !unicode.IsSpace(s) {
			flg = true
		}
	}

	return flg
}