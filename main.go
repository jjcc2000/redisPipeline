package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("TCP Listen Error", err)
	}

	log.Printf("Listener: %v", listener)
	// for {
	// 	conn, err := ln.Accept()
	// 	if err != nil {

	// 	}

	// }
}
