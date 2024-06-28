package repository

import (
	"errors"
	"user/internal/service"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID         uint   `json:"user_id" gorm:"primary_key"` // 用户ID
	UserName       string `json:"username" gorm:"unique"`     // 用户名
	NickName       string `json:"nickname"`                   // 昵称
	PasswordDigest string `json:"password_digest"`            // 加密后的密码
}

const (
	PasswordCost = 12 // 密码加密难度
)

// CheckUserExist 检查用户是否存在
func (u *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("username =?", req.UserName).First(&u).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// ShowUserInfo 获取用户信息
func (u *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := u.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("UserName Not Exist")
}

// CreateUser 创建用户
func (*User) UserCreate(req *service.UserRequest) (err error) {
	var count int64
	DB.Where("username =?", req.UserName).Count(&count)
	if count > 0 {
		return errors.New("UserName Already Exist")
	}
	user := User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	// 加密密码
	_ = user.SetPassword(req.PassWord)
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// SetPassword 加密密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

// BuildUser 序列化用户信息
func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		UserID:   uint32(item.UserID),
		UserName: item.UserName,
		NickName: item.NickName,
	}
	return &userModel
}
