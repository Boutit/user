package server

import (
	"context"
	"fmt"

	"github.com/Boutit/user/api"
)

func (s userServiceServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	err := s.userStore.CreateUser(ctx, req.User)
	if err != nil {
		return nil, err
	}
	fmt.Println(req.User)
	return &api.CreateUserResponse{}, nil
}