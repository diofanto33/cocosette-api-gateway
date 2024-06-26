package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// importing generated stubs
	gen "github.com/diofanto33/cocosette-api-gateway/hello"

	"google.golang.org/grpc"
)

// GreeterServerImpl will implement the service defined in protocol buffer definitions
type GreeterServerImpl struct {
	gen.UnimplementedGreeterServer
}

// SayHello is the implementation of RPC call defined in protocol definitions.
// This will take HelloRequest message and return HelloReply
func (g *GreeterServerImpl) SayHello(ctx context.Context, request *gen.HelloRequest) (*gen.HelloReply, error) {
	return &gen.HelloReply{
		Message: fmt.Sprintf("hello there %s", request.Name),
	}, nil
}
func main() {
	// create new gRPC server
	server := grpc.NewServer()
	// register the GreeterServerImpl on the gRPC server
	gen.RegisterGreeterServer(server, &GreeterServerImpl{})
	// start listening on port :8080 for a tcp connection
	if l, err := net.Listen("tcp", ":8080"); err != nil {
		log.Fatal("error in listening on port :8080", err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
