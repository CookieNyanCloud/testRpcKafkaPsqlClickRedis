run:
	go run cmd/app/main.go

protos:
	protoc --go_out=./protos --go_opt=paths=source_relative \
        --go-grpc_out=./protos --go-grpc_opt=paths=source_relative \
        protos/users.proto

up:
	docker-compose up -d --build
.PHONY: protos run