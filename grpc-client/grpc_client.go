package main

import (
	"flag"
	"fmt"
	"log"

	"google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
)

func main() {
	ClientConnection("")
}

func ClientConnection(server_address string) {
	flag.Parse()
	var grpc_opts []grpc.DialOption
	grpc_opts = append(grpc_opts, grpc.WithInsecure())
	grpc_connection, err := grpc.Dial(server_address, grpc_opts)
	if err != nil {
		log.Fatalf("Failed to Dial #{err}")
	}
	if grpc_connection != nil {
		fmt.Printf("Connection Succeeded to %v", server_address)
	}
}
