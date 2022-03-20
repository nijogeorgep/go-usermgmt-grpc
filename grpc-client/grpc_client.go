package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pb "go-usermgmt-grpc/usermgmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	server_address = "0.0.0.0:50051"
)

func main() {

	grpc_credentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("Couldn't load TLS credentials: %v", err)
	}
	//var grpc_opts []grpc.DialOption
	//grpc_opts = append(grpc_opts, grpc.WithInsecure())
	//grpc_opts = append(grpc_opts, grpc.WithBlock())

	grpc_connection, err := grpc.Dial(server_address, grpc.WithTransportCredentials(grpc_credentials))
	if err != nil {
		log.Fatalf("Failed to Dial %v", err)
	}
	if grpc_connection != nil {
		fmt.Printf("Connection Succeeded to %v ", server_address)
	}
	defer grpc_connection.Close()

	grpc_client := pb.NewUserManagementClient(grpc_connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["NIJO"] = 33
	new_users["GEORGE"] = 60

	for name, age := range new_users {
		response, err := grpc_client.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("Couldn't create User %v", err)
		}
		log.Printf(`User Details : 
		Name : %s
		Age : %d
		Id : %d
		`, response.GetName(), response.GetAge(), response.GetId())
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("certificates/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
  
	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("certificates/client-cert.pem", "certificates/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
