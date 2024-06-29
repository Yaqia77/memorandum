package handler

import (
	"context"
	"sync"

	"github.com/yaqia77/memorandum/apps/user/internal/repository/db/dao"
	"github.com/yaqia77/memorandum/apps/user/internal/service"
	"github.com/yaqia77/memorandum/pkg/e"
)

type UserService struct {
	service.UserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
	service.UnimplementedUserServiceServer
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success
	r, err := dao.NewUserDao(ctx).GetUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		return
	}
	resp.UserDetail = &service.UserResponse{
		UserId:   r.UserID,
		UserName: r.UserName,
		NickName: r.UserName,
	}
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserCommonResponse, err error) {
	resp = new(service.UserCommonResponse)
	resp.Code = e.Success
	err = dao.NewUserDao(ctx).CreateUser(req)
	if err != nil {
		resp.Code = e.Error
		return
	}
	resp.Data = e.GetMsg(uint(resp.Code))
	return
}

func (u *UserSrv) UserLogout(ctx context.Context, request *service.UserRequest) (resp *service.UserCommonResponse, err error) {
	return
}
