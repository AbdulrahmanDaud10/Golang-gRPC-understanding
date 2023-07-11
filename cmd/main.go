package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	_, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer()
	fmt.Println("---------", gRPCServer)

}
