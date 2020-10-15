package main

import (
	"io"
	"log"
	"net"
	"os"
	"flag"
)

func main() {
	var NewWork, Tokyo, London string
	flag.StringVar(&NewWork, "NewWork", "localhost:8010", "area 1")
	flag.StringVar(&Tokyo, "Tokyo", "localhost:8020", "area 2")
	flag.StringVar(&London, "London", "localhost:8030", "area 3")

	var area [3]string{NewWork}
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
