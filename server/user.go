package main

import (
	"context"
	"fmt"
	"net"

	"github.com/thejasbabu/protobuf-example/user"
	"google.golang.org/grpc"
)

type UserHandler struct {
	users []user.UserInfo
}

func (h UserHandler) SearchUserBy(email string) (user.UserInfo, error) {
	var selectedUser user.UserInfo
	var found bool
	for _, u := range h.users {
		if u.Email == email {
			selectedUser = u
			found = true
			break
		}
	}
	if found {
		return selectedUser, nil
	}
	return selectedUser, fmt.Errorf("user with email %s not found", email)
}

func (h *UserHandler) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.UserInfo, error) {
	email := req.GetEmail()
	fmt.Printf("Recieved search request for user %s", email)
	selectedUser, err := h.SearchUserBy(email)
	if err != nil {
		return nil, err
	}

	resp := user.UserInfo{
		Email: selectedUser.GetEmail(),
		Name:  selectedUser.GetName(),
		Phone: selectedUser.GetPhone(),
		Address: &user.UserInfo_Address{Street: selectedUser.GetAddress().GetStreet(),
			City:  selectedUser.GetAddress().GetCity(),
			State: selectedUser.GetAddress().GetState(),
			Zip:   selectedUser.GetAddress().GetZip(),
		},
	}
	return &resp, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *user.UserInfo) (*user.UserStatus, error) {
	email := req.GetEmail()
	fmt.Printf("Recieved creation request for user %s", email)
	_, err := h.SearchUserBy(email)
	if err != nil {
		h.users = append(h.users, *req)
		return &user.UserStatus{Error: ""}, nil
	}
	return &user.UserStatus{Error: "User already present"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Printf("error starting grpc server: %s", err.Error())
		panic(1)
	}

	server := grpc.NewServer()
	userHandler := UserHandler{users: []user.UserInfo{}}
	user.RegisterUserServer(server, &userHandler)

	if err := server.Serve(lis); err != nil {
		fmt.Printf("error: %s", err.Error())
		panic(1)
	}
}
