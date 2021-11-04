// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"grpcDemo/api"

	"google.golang.org/grpc"
)

const (
	// serverAddr rpc 服务端地址
	serverAddr  = "localhost:8081"
	defaultName = "dounine"
)

var (
	logger = log.New(os.Stdout, "[client] ", log.Lshortfile|log.Ldate|log.Ltime)
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create the client of HelloService.
	c := api.NewHelloServiceClient(conn)

	// Contact the server, call service method SayHello() and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &api.HelloRequest{Name: name, Age: "18"})
	if err != nil {
		logger.Fatalf("could not greet: %v", err)
	}
	logger.Printf("Response: %s", r.GetMessage())
}
