package main

import (
	"context"
	"fmt"
	"time"

	"github.com/thejasbabu/protobuf-example/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error connecting to sever: %s", err.Error())
		panic(1)
	}

	defer conn.Close()
	client := user.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userInfo := user.UserInfo{Email: "user1@test.com", Name: "user1"}
	resp, err := client.CreateUser(ctx, &userInfo)
	if err != nil {
		fmt.Printf("error from server: %s", err.Error())
		panic(1)
	}
	if resp.GetError() != "" {
		fmt.Printf("error from server: %s", resp.GetError())
		panic(1)
	}
	userRequest := user.GetUserRequest{Email: "user1@test.com"}
	userData, err := client.GetUser(ctx, &userRequest)
	if err != nil {
		fmt.Printf("error from server: %s", err.Error())
		panic(1)
	}
	fmt.Printf("user name: %s, user email: %s", userData.GetName(), userData.GetEmail())
	invalidUserRequest := user.GetUserRequest{Email: "user2@test.com"}
	_, err = client.GetUser(ctx, &invalidUserRequest)
	if err != nil {
		fmt.Printf("error from server: %s", err.Error())
		panic(1)
	}
}
