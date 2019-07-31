package main

import (
	"fmt"
)

func main() {
	s := []string{"0", "1", "2", "2", "2", "3", "3", "4", "5"}
	fmt.Println(removeDuplicate(s))
}

func removeDuplicate(s []string) []string {
	var ans []string
	for i := 0; i < len(s) -1 ; {
		if s[i] != s[i+1] {
			ans = append(ans, s[i])
		}
		i++
	}
	return ans
}