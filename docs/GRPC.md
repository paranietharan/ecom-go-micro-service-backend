# gRPC Service Guide

This document provides detailed information about the gRPC service in the E-Commerce Microservice Backend.

## Overview

The project uses [gRPC](https://grpc.io/) for communication between the API service and the core business service. gRPC utilizes Protocol Buffers as the interface definition language and serialization format.

## Service Definition

The service interface is defined in the `ecom-grpc/pb/api.proto` file:

```protobuf
syntax = "proto3";

package pb;

option go_package = "ecom-go-micro-service-backend/ecom-grpc/pb";

import "google/protobuf/timestamp.proto";

// Message definitions for Products, Orders, Users, etc.
// ...

service ecomm {
  // Product operations
  rpc CreateProduct(ProductReq) returns (ProductRes) {}
  rpc GetProduct(ProductReq) returns (ProductRes) {}
  rpc ListProducts(ProductReq) returns (ListProductRes) {}
  rpc UpdateProduct(ProductReq) returns (ProductRes) {}
  rpc DeleteProduct(ProductReq) returns (ProductRes) {}

  // Order operations
  rpc CreateOrder(OrderReq) returns (OrderRes) {}
  rpc GetOrder(OrderReq) returns (OrderRes) {}
  rpc ListOrders(OrderReq) returns (ListOrderRes) {}
  rpc DeleteOrder(OrderReq) returns (OrderRes) {}

  // User operations
  rpc CreateUser(UserReq) returns (UserRes) {}
  rpc GetUser(UserReq) returns (UserRes) {}
  rpc ListUsers(UserReq) returns (ListUserRes) {}
  rpc UpdateUser(UserReq) returns (UserRes) {}
  rpc DeleteUser(UserReq) returns (UserRes) {}

  // Session operations
  rpc CreateSession(SessionReq) returns (SessionRes) {}
  rpc GetSession(SessionReq) returns (SessionRes) {}
  rpc RevokeSession(SessionReq) returns (SessionRes) {}
  rpc DeleteSession(SessionReq) returns (SessionRes) {}
}
```

## Message Types

### Product Messages

```protobuf
message ProductReq {
  int64 id = 1;
  string name = 2;
  string image = 3;
  string category = 4;
  string description = 5;
  int64 rating = 6;
  int64 num_reviews = 7;
  float price = 8;
  int64 count_in_stock = 9;
}

message ProductRes {
  int64 id = 1;
  string name = 2;
  string image = 3;
  string category = 4;
  string description = 5;
  int64 rating = 6;
  int64 num_reviews = 7;
  float price = 8;
  int64 count_in_stock = 9;
  google.protobuf.Timestamp created_at = 10;
  google.protobuf.Timestamp updated_at = 11;
}

message ListProductRes {
  repeated ProductRes products = 1;
}
```

### Order Messages

```protobuf
message OrderItem {
  string name = 1;
  int64 quantity = 2;
  string image = 3;
  float price = 4;
  int64 product_id = 5;
}

message OrderReq {
  int64 id = 1;
  repeated OrderItem items = 2;
  string payment_method = 3;
  float tax_price = 4;
  float shipping_price = 5;
  float total_price = 6;
  int64 user_id = 7;
}

message OrderRes {
  int64 id = 1;
  repeated OrderItem items = 2;
  string payment_method = 3;
  float tax_price = 4;
  float shipping_price = 5;
  float total_price = 6;
  int64 user_id = 7;
  google.protobuf.Timestamp created_at = 8;
  google.protobuf.Timestamp updated_at = 9;
}

message ListOrderRes {
  repeated OrderRes orders = 1;
}
```

### User Messages

```protobuf
message UserReq {
  int64 id = 1;
  string name = 2;
  string email = 3;
  string password = 4;
  bool is_admin = 5;
}

message UserRes {
  int64 id = 1;
  string name = 2;
  string email = 3;
  string password = 4;
  bool is_admin = 5;
  google.protobuf.Timestamp created_at = 6;
}

message ListUserRes {
  repeated UserRes users = 1;
}
```

### Session Messages

```protobuf
message SessionReq {
  string id = 1;
  string user_email = 2;
  string refresh_token = 3;
  bool is_revoked = 4;
  google.protobuf.Timestamp expires_at = 5;
}

message SessionRes {
  string id = 1;
  string user_email = 2;
  string refresh_token = 3;
  bool is_revoked = 4;
  google.protobuf.Timestamp expires_at = 5;
}
```

## Generating gRPC Code

After modifying the `.proto` file, you need to regenerate the Go code using the Protocol Buffers compiler:

1. **Install the Protocol Buffers Compiler**

   Follow the instructions at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/) to install the `protoc` compiler for your platform.

2. **Install Go Plugins for Protocol Buffers**

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. **Generate Go Code**

   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       ecom-grpc/pb/api.proto
   ```

   This will generate two files:
   - `api.pb.go`: Contains the protocol message definitions
   - `api_grpc.pb.go`: Contains the gRPC service definitions

## Implementing the gRPC Server

The gRPC server implementation is in the `ecom-grpc/server/server.go` file. This is where the business logic for all service methods is implemented.

A service method implementation typically:

1. Validates inputs
2. Calls the database storer functions
3. Transforms data if needed
4. Returns the response

## Using the gRPC Client

The API service acts as a gRPC client that calls the gRPC server methods. The client is created in the API service code:

```go
conn, err := grpc.Dial(grpcSvcAddr, grpc.WithInsecure())
if err != nil {
    log.Fatalf("Failed to connect to gRPC server: %v", err)
}
defer conn.Close()

client := pb.NewEcommClient(conn)
```

## Best Practices

1. **Use Well-Defined Message Types**: Keep message definitions clean and follow Protocol Buffers best practices.

2. **Versioning**: Consider versioning your API if making breaking changes (e.g., `v1/api.proto`, `v2/api.proto`).

3. **Error Handling**: Use appropriate gRPC error codes and include descriptive error messages.

4. **Streaming**: Consider using streaming RPCs for large datasets or real-time updates.

5. **Authentication**: Implement authentication using interceptors.

## Troubleshooting

### Common Issues

1. **Mismatched Versions**: If you get errors about missing methods or fields, ensure both services are using the same version of the generated code.

2. **Connection Issues**: Check that the gRPC service is running and accessible on the specified address and port.

3. **Compilation Errors**: Make sure your `.proto` file is valid and all imported types are defined.

4. **Context Timeouts**: Set appropriate timeouts for gRPC client calls based on expected operation time.

### Debugging Tools

- [grpcurl](https://github.com/fullstorydev/grpcurl): Command-line tool for interacting with gRPC servers
- [BloomRPC](https://github.com/uw-labs/bloomrpc): GUI client for testing gRPC services
