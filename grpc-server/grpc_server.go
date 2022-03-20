package main

import (
	"context"
	"crypto/tls"
	"log"
	"math/rand"
	"net"

	pb "go-usermgmt-grpc/usermgmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	credentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("Couldn't load TLS Credentials : %v", err)
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", port)
	}

	grpc_server := grpc.NewServer(grpc.Creds(credentials))

	pb.RegisterUserManagementServer(grpc_server, &UserManagementServer{})
	log.Printf("Server Listening at %v", listener.Addr())

	if err := grpc_server.Serve(listener); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("certificates/server-cert.pem", "certificates/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
