package main

// This file exposes a gRPC server that listens on port 8080 for a method called DataCallback.

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/goncalo-leal/go-fixture/proto/data"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedDataServiceServer
}

// DataCallback implements helloworld.GreeterServer
func (s *server) DataCallback(ctx context.Context, req *pb.DataReceived) (*pb.DataResponse, error) {
	log.Printf("Received: %v", req.Data)

	return &pb.DataResponse{Status: "Success"}, nil
}

func main() {
	fmt.Println("1")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen on port 8080: %v", err)
	}

	fmt.Println("2")
	s := grpc.NewServer()
	pb.RegisterDataServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
