package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "sha256" {
		fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[1])))
	} else if os.Args[1] == "sha384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(os.Args[1])))
	} else if os.Args[1] == "sha512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(os.Args[1])))
	} else {
		fmt.Println("Cannot use first argument without sha256, sha384 or sha512")
	}
}
