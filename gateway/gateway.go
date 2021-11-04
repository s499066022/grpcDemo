// package main provide http gateway according to rpc-server
package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "grpcDemo/api"
)

const (
	port = ":8080"
)

var (
	// serverAddr gRPC server addr
	serverAddr = flag.String("server_addr", "localhost:8081", "address of YourServer")
	logger     = log.New(os.Stdout, "[gateway] ", log.Lshortfile|log.Ldate|log.Ltime)
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register HelloService handler.
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, *serverAddr, opts)
	if err != nil {
		return err
	}

	// Listening the port and open the serve.
	logger.Print("gateway is running in ", port)
	return http.ListenAndServe(port, mux)
}

func main() {
	flag.Parse()

	err := run()
	if err != nil {
		logger.Fatal(err)
	}
}
