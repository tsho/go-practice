package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8000, "port number")
}

// flag.Int("port", 8080, "listen port")

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()

	server := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening at localhost:%d\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}
