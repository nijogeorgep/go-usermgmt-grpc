package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "go-usermgmt-grpc/usermgmt"

	"google.golang.org/grpc"
)

const (
	server_address = "localhost:50051"
)

func main() {

	//var grpc_opts []grpc.DialOption
	//grpc_opts = append(grpc_opts, grpc.WithInsecure())
	//grpc_opts = append(grpc_opts, grpc.WithBlock())

	grpc_connection, err := grpc.Dial(server_address, grpc.WithInsecure(), grpc.WithBlock())
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
