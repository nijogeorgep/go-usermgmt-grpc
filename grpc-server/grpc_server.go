package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	StartServer(4000)
}

func StartServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("#{port}"))
	if err != nil {
		log.Fatalf("Failed to Listen: #{err}")
	}
	grpc_server := grpc.NewServer()

	if err = grpc_server.Serve(listener); err != nil {
		log.Fatalf("Unable to start gRPC Server: #{err}")
	}
}
