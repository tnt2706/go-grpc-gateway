.PHONY:	install gen-grpc server

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get github.com/grpc-ecosystem/grpc-gateway/v2@v2.16.0


gen-grpc:
	protoc --go_out=pkg/generate --go_opt=Mprotos/*.proto=pb \
    --go-grpc_out=pkg/generate --go-grpc_opt=Mprotos/*.proto=pb \
    internal/proto/*.proto

server:
	go run cmd/main.go