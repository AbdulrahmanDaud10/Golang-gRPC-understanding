package main

import (
	"log"
	"net"
)

func main() {
	_, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}
