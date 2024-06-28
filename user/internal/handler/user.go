package handler

import (
	"context"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/e"
)

type UserService struct {
	service.UserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// UserLogin 用户登录
func (u *UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success

	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	//序列化
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

// UserRegister 用户注册
func (u *UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success

	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	//序列化
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil

}

// mustEmbedUnimplementedUserServiceServer 是 gRPC 自动生成的方法，用于确保接口完整实现
func (u *UserService) mustEmbedUnimplementedUserServiceServer() {}
