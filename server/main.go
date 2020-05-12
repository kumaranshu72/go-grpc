package main

import (
	"fmt"
	"grpc/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

// starts a grpc server and waits for connection
func main() {
	// create a listner on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	// creating a server instance
	s := api.Server{}

	// creating a grpc object
	grpcServer := grpc.NewServer()

	// attach the ping service to the server
	api.RegisterPingServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

