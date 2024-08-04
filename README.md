# go-grpc-gateway

## Overview

gRPC-Gateway helps you to provide your APIs in both gRPC and RESTful style at the same time.

![alt text](./images/architecture_introduction_diagram.svg)

## Document

- <https://github.com/grpc-ecosystem/grpc-gateway>
- <https://grpc-ecosystem.github.io/grpc-gateway/>

### Generate

- **Use default**

```proto
protoc --go_out=pkg/proto --go_opt=paths=source_relative \
--go-grpc_out=pkg/proto --go-grpc_opt=paths=source_relative \
internal/proto/calculator_service.proto
```

![alt text](./images/generate_default.png)

- **Use custom**

```proto
option go_package = "./pb";

protoc --go_out=pkg/generate --go_opt=Mprotos/*.proto=pb \
--go-grpc_out=pkg/generate --go-grpc_opt=Mprotos/*.proto=pb \
internal/proto/*.proto
```

![alt text](./images/generate_custom.png)
