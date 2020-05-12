package main

import (
	"fmt"
	"grpc/api"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	// craete TSL credentials
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	// creating a grpc object
	grpcServer := grpc.NewServer(opts...)

	// attach the ping service to the server
	api.RegisterPingServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
