package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/diofanto33/cocosette-proto/golang/auth"
)

func main() {
	// Create a new gRPC client connection to the AuthService
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create a new gRPC-Gateway mux
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	// Register the AuthService with the gRPC-Gateway mux
	err = auth.RegisterAuthServiceHandler(context.Background(), mux, conn)
	if err != nil {
		panic(err)
	}

	// Create a new HTTP server that uses the gRPC-Gateway mux
	http.ListenAndServe(":8080", mux)
}
