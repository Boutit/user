package server

import (
	"context"

	"github.com/Boutit/user/api"
	"github.com/Boutit/user/internal/models"
	"github.com/Boutit/user/internal/utils"
)

func (s userServiceServer) SignupUser(ctx context.Context, req *api.SignupUserRequest) (*api.SignupUserResponse, error) {
	var p, e string
	if utils.IsEmail(req.EmailOrPhone){
		e = req.EmailOrPhone
	} else if utils.IsPhoneNumber(req.EmailOrPhone){
		p = req.EmailOrPhone
	}
	h := utils.HashPassword(req.Password)
	user := &models.User{
		Email: e,
		Name: req.Name,
		Password: h,
		Phone: p,
		Username: req.Username,
	}

	err := s.userStore.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &api.SignupUserResponse{}, nil
}

func (s userServiceServer) LoginUser(ctx context.Context, req *api.LoginUserRequest) (*api.LoginUserResponse, error) {
	var p string
	if utils.IsEmail(req.EmailPhoneOrUsername){
		 // get user
	} else if utils.IsPhoneNumber(req.EmailPhoneOrUsername){
		//  get user
	}
	// compare passwords
	utils.ComparePassword(p, req.Password)
	return &api.LoginUserResponse{}, nil
}