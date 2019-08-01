package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"strings"
	"io"
	"log"
)

func exit() {
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

type ResponseResultURL struct {
	Poster string
	Number int
}

func main() {
	if len(os.Args) < 3 {
		exit()
	}
	
	apikey := os.Args[1]
	name := strings.Join(os.Args[2:]," ")
	
	resp, err := http.Get("http://omdbapi.com/?t="+name+"&apikey="+apikey)
	if err != nil {
		log.Fatal(err)
	}
	
	var result ResponseResultURL
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	resp.Body.Close()
	
	resp, err = http.Get(result.Poster)
	if err != nil {
		log.Fatal(err)
	}
	//open file for writing
	image, err := os.Create("./" + name + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	//io.Copy is good for a huge files
	size, err := io.Copy(image, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	resp.Body.Close()
	image.Close()
	fmt.Printf("%s with %v bytes downloaded\n", name, size)
}
