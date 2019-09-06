package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		if strings.HasPrefix(word, "$") {
			words[i] = f(word[1:])
		}
	}
	return strings.Join(words, " ")
}


func main() {
	text :=  "$hello, $world"

	f := func(text string) string {
		return text
	}
	
	fmt.Println(expand(text, f))
}
