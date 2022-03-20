package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "go-usermgmt-grpc/usermgmt"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received : %v and %d", in.GetName(), in.GetAge())
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", port)
	}

	grpc_server := grpc.NewServer()

	pb.RegisterUserManagementServer(grpc_server, &UserManagementServer{})
	log.Printf("Server Listening at %v", listener.Addr())

	if err := grpc_server.Serve(listener); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
