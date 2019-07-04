package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"

	"../ex01/tempconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			targ := scanner.Text()
			convert(targ)
		}
	} else {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	}
}
	

func convert(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}