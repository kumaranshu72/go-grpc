package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server respresents the grpc server
type Server struct {
}

// SayHello respond to ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Recieving message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
