.PHONY:	install gen-grpc server document document2

install:
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

gen-grpc:
		protoc -I./ \
        --go_out=${APP_HOME}/pkg \
        --go-grpc_out=${APP_HOME}/pkg \
        --go-grpc_opt=require_unimplemented_servers=false \
      	google/api/*.proto \
        app/if/*.proto \
        app/model/*.proto \
        common/*.proto \
        service/calculator_service.proto \
        service/sample_service.proto

gen-grpc-gateway:
		protoc -I./ \
			-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.16.0 \
			-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
			--go_out=${APP_HOME}/pkg \
			--go-grpc_out=${APP_HOME}/pkg \
			--go-grpc_opt=require_unimplemented_servers=false \
		 	google/api/*.proto \
            app/if/*.proto \
            app/model/*.proto \
            common/*.proto \
            service/calculator_service.proto \
            service/sample_service.proto

		protoc -I./ \
			  -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.16.0 \
			  -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
			  --grpc-gateway_out $(APP_HOME)/pkg \
			  --grpc-gateway_opt logtostderr=true \
			  --openapiv2_out=$(APP_HOME)/pkg \
			  --openapiv2_opt=logtostderr=true \
			  service/calculator_service.proto \
			  service/sample_service.proto


document:
		protoc -I./ \
			-I$(APP_HOME)/internal/pb/include \
        	-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.16.0 \
        	-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
			--openapiv2_out ${APP_HOME}/pkg \
			--openapiv2_opt=logtostderr=true \
            service/calculator_service.proto \
            service/sample_service.proto

document2:
	protoc -I./ -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --openapiv2_out=./doc --doc_out=./doc --doc_opt=html,index.html app.proto

server:
	go run cmd/main.go
