IMAGE_NAME := go-usermgmt-grpc
IMAGE_TAG := latest

proto:
	protoc -I . \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	usermgmt/*.proto

server:
	go build -o bin/grpc_server grpc-server/grpc_server.go
	@chmod +x bin/grpc_server

client:
	go build -o bin/grpc_client grpc-client/grpc_client.go
	@chmod +x bin/grpc_client

image:
	@docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

run-bash:
	@docker run -i -t $(IMAGE_NAME):$(IMAGE_TAG) /bin/bash

run:
	@docker run -p 4000:4000 -it $(IMAGE_NAME):$(IMAGE_TAG)
