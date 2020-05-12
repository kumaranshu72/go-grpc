package main

import (
	"grpc/api"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// main function to create client call
func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	// defer conn.close()

	c := api.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "Foo"})

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Greeting)
}
