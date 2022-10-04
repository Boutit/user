package server

import (
	"github.com/Boutit/user/api"
	"github.com/Boutit/user/internal/store"
)

type UserServiceServer interface {
	api.UserServiceServer
}

type userServiceServer struct {
	api.UnimplementedUserServiceServer
	userStore store.UserStore
}

func NewUserServiceServer(store store.UserStore) UserServiceServer {
	return &userServiceServer{
		userStore: store,
	}
}