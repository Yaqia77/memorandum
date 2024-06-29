package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/yaqia77/memorandum/apps/user/internal/repository/db/model"
	"github.com/yaqia77/memorandum/apps/user/internal/service"
	userPb "github.com/yaqia77/memorandum/apps/user/internal/service"
	"github.com/yaqia77/memorandum/pkg/util/logger"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// CheckUserExist 检查用户是否存在
func (dao *UserDao) CheckUserExist(req *service.UserRequest) bool {
	if err := dao.Where("username =?", req.UserName).First(&model.User{}).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// GetUserInfo 获取用户信息
func (dao *UserDao) GetUserInfo(req *userPb.UserRequest) (r *model.User, err error) {
	err = dao.Model(&model.User{}).Where("user_name=?", req.UserName).
		First(&r).Error

	return
}

// CreateUser 用户创建
func (dao *UserDao) CreateUser(req *userPb.UserRequest) (err error) {
	var user model.User
	var count int64
	dao.Model(&model.User{}).Where("user_name = ?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("UserName Exist")
	}

	user = model.User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	_ = user.SetPassword(req.Password)
	if err = dao.Model(&model.User{}).Create(&user).Error; err != nil {
		logger.LogrusObj.Error("Insert User Error:" + err.Error())
		return
	}

	return
}

// BuildUser 序列化用户信息
func BuildUser(item model.User) *service.UserResponse {
	userModel := service.UserResponse{
		UserId:   item.UserID,
		UserName: item.UserName,
		NickName: item.NickName,
	}
	return &userModel
}
