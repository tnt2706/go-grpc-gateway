.PHONY:	install gen-grpc server document clean

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
        app/if/*.proto \
        app/model/*.proto \
        common/*.proto \
        service/calculator_service.proto \

gen-grpc-gateway:
		protoc -I./ \
			-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.16.0 \
			-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
			--go_out=${APP_HOME}/pkg \
			--go-grpc_out=${APP_HOME}/pkg \
			--go-grpc_opt=require_unimplemented_servers=false \
            app/if/*.proto \
            app/model/*.proto \
            common/*.proto \
            service/calculator_service.proto \

		protoc -I./ \
			  -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.16.0 \
			  -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
			  --grpc-gateway_out $(APP_HOME)/pkg \
			  --grpc-gateway_opt logtostderr=true \
			  service/calculator_service.proto \


document:
		protoc -I./ \
        	-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.16.0 \
        	-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
			--openapiv2_out ${APP_HOME}/docs \
			--openapiv2_opt=logtostderr=true \
            service/calculator_service.proto

clean:
	rm -rf ${APP_HOME}/pkg/pb/*
	rm -rf ${APP_HOME}/docs/*
