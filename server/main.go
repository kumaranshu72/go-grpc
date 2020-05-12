package main

import (
	"context"
	"fmt"
	"grpc/api"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// private type for context key
type contextKey int

const (
	clientIdKey contextKey = iota
)

// authenticating the client
func authenticateClient(ctx context.Context, s *api.Server) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		clientLogin := strings.Join(md["login"], "")
		clientPassword := strings.Join(md["password"], "")

		if clientLogin != "anshu" {
			return "", fmt.Errorf("unknown user %s", clientLogin)
		}
		if clientPassword != "anshu" {
			return "", fmt.Errorf("bad password %s", clientPassword)
		}
		log.Printf("authenticated client: %s", clientLogin)
		return "42", nil
	}

	return "", fmt.Errorf("missing credentials")
}

// unaryInterceptor calls authenticateClient with current context
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	s, ok := info.Server.(*api.Server)
	if !ok {
		return nil, fmt.Errorf("unable to cast server")
	}
	clientID, err := authenticateClient(ctx, s)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, clientIdKey, clientID)
	return handler(ctx, req)
}

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
	opts := []grpc.ServerOption{grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor)}

	// creating a grpc object
	grpcServer := grpc.NewServer(opts...)

	// attach the ping service to the server
	api.RegisterPingServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
