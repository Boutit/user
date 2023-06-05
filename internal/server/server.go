package server

import (
	"github.com/Boutit/user/api/protos/boutit/user"
	"github.com/Boutit/user/internal/store"
)

type UserServiceServer interface {
	user.UserServiceServer
}

type userServiceServer struct {
	user.UnimplementedUserServiceServer
	userStore store.UserStore
}

func NewUserServiceServer(store store.UserStore) UserServiceServer {
	return &userServiceServer{
		userStore: store,
	}
}