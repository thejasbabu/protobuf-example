# Golang Protobuf example with grpc

A simple grpc server and client using protobuf to communicate. Consists of Get and Create methods.

## Set-up

1. Install the protobuf compiler from [here](https://github.com/protocolbuffers/protobuf/blob/master/README.md#protocol-compiler-installation)

2. Compile the proto file to generate the go clients and services

      ```
      protoc -I user/ user/user.proto --go_out=plugins=grpc:user
      ```
3. Compile both server and client code

    ```
    go build -o user_client client/user.go
    go build -o user_server server/user.go
    ```
4. Run both the client and server.

