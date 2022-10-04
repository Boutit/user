package server

import (
	"context"

	"github.com/Boutit/user/api"
)

func (s userServiceServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	err := s.userStore.CreateUser(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &api.CreateUserResponse{}, nil
}