package main

import (
	"log"
	"net"

	pb "github.com/AbdulrahmanDaud10/Golang-gRPC-understanding/proto"
	"github.com/AbdulrahmanDaud10/Golang-gRPC-understanding/services"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer()
	service := &services.UserServiceServer{}

	pb.RegisterUserServiceServer(gRPCServer, service)
	err = gRPCServer.Serve(listen)

	if err != nil {
		log.Fatalf("Error strating server: %v", err)
	}

}
