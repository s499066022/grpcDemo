// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"os"

	"grpcDemo/api"

	"google.golang.org/grpc"
)

const (
	port = ":8081"
)

var (
	logger = log.New(os.Stdout, "[server] ", log.Lshortfile|log.Ldate|log.Ltime)
)

// server is used to implement api.HelloServiceServer.
type server struct {
	api.UnimplementedHelloServiceServer
}

// SayHello implements api.SayHello
func (s *server) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloReply, error) {
	logger.Print("Received: ", req.String())
	return &api.HelloReply{Message: "Hello " + req.GetName() + ", your age is " + req.GetAge()}, nil
}

func main() {
	// listening the tcp port
	listen, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// register the server
	s := grpc.NewServer()
	api.RegisterHelloServiceServer(s, &server{})

	// open the serve
	logger.Print("server starting in ", port)
	err = s.Serve(listen)
	if err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
