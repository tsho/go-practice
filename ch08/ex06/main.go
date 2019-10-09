package main

import (
	"flag"
	"fmt"
	"log"

	"../../gopl.io/ch5/links"
)

var depth = flag.Int("depth", 1, "depth links")

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	flag.Parse()
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	var n, d int                     // number of pending sends to worklist

	n++
	counter := make([]int, *depth+2)
	counter[d] = n

	go func() {
		worklist <- flag.Args()
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist

		if d > *depth {
			continue
		}
		for _, link := range list {
			if !seen[link] {
				n++ // counter++
				counter[d+1]++

				seen[link] = true
				unseenLinks <- link
			}
		}
		if counter[d]--; counter[d] == 0 {
			d++
		}
	}
	close(unseenLinks)
}
