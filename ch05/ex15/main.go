package main

import "fmt"

//!+
func max(vals ...int) int {
	if len(vals) == 0 {
		fmt.Println("set args above one argument")
		return 0
	}

	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) == 0 {
		fmt.Println("set args above one argument")
		return 0
	}

	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	fmt.Println(max())           //  "0"
	fmt.Println(max(3))          //  "3"
	fmt.Println(max(1, 2, 3, 4)) //  "10"

	values := []int{1, 2, 3, 4}
	fmt.Println(max(values...)) // "10"

	fmt.Println(min())           //  "0"
	fmt.Println(min(3))          //  "3"
	fmt.Println(min(1, 2, 3, 4)) //  "10"

	fmt.Println(min(values...)) // "10"
}