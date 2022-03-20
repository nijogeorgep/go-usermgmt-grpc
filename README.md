### This is a small project to demonstrate gRPC with Golang

Steps to Run this Project in a local machine

- Install Go
- Install ProtoBuff
- Install protocol compiler plugins for Go  
- Generate SSL Certificates using `make certificate`
- Generate Protobuff Go Files using `make proto`
- Create Server Executable using `make server`
- Create Client Executable using `make client`
- Run `grpc_server` first from `BIN` directory
- Run `grpc_client` from `BIN` directory

#### Generate a Private Key and a CSR
```
openssl req \
       -newkey rsa:2048 -nodes -keyout grpc.key \
       -out grpc.csr
```
#### Generate a Self-Signed Certificate from an Existing Private Key and CSR
```
openssl x509 \
       -signkey grpc.key \
       -in grpc.csr \
       -req -days 365 -out grpc.crt
```