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
    internal/pb/service/*.proto

server:
	go run cmd/main.go


gen-grpc-gateway:
	protoc -I . --grpc-gateway_out pkg/generate \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    internal/proto/*.proto