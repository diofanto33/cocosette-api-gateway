package main

import (
	"context"
	"log"
	"net/http"

	"github.com/diofanto33/cocosette-proto/golang/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// create a client connection to the gRPC server we just started
	// this is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
		return
	}

	defer conn.Close()

	// muxserver object
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption("application/json", &runtime.JSONPb{}),
	)
	// register the Auth service handler with the client connection and mux
	err = auth.RegisterAuthServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
		return
	}

	// create a new HTTP server and pass the gwmux handler to it
	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	// start the HTTP Server
	log.Println("Starting HTTP server on", gwServer.Addr)
	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

}
