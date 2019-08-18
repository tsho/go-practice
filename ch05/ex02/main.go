// Exercise 5.2

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	cnt_map := make(map[string]int)
	visit(cnt_map, doc)
	fmt.Println(cnt_map)
}

// visit appends to links each link found in n and returns the result.
func visit(cnt_map map[string]int, n *html.Node)  {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		cnt_map[n.Data]++
	}
	visit(cnt_map, n.FirstChild)
	visit(cnt_map, n.NextSibling)
}
