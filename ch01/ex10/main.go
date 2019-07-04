// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"log"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file, err := os.Create(`./output.txt`) 
	if err != nil {
        log.Fatal(err)
    }
	defer file.Close()  //全部の
	
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
		time.Sleep(3 * time.Second)
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		file.Write(([]byte)(<-ch))

		file.Write(([]byte)(<-ch))
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url, resp.Body)
	resp.Body.Close() // don't leak resources
}

//!-