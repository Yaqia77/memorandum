package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util/jwt"
	"context"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	// gin.Key 中获取服务实例
	userService := ctx.Keys["user"].(service.UserServiceClient)

	userResq, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	r := res.Response{
		Data:   userResq,
		Status: uint(userResq.Code),
		Msg:    e.GetMsg(uint(userResq.Code)),
		Error:  err.Error(),
	}
	ctx.JSON(200, r)
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	// gin.Key 中获取服务实例
	userService := ctx.Keys["user"].(service.UserServiceClient)

	userResq, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := jwt.GenerateToken(userResq.UserDetail.UserID)
	r := res.Response{
		Data:   res.TokenData{User: userResq.UserDetail, Token: token},
		Status: uint(userResq.Code),
		Msg:    e.GetMsg(uint(userResq.Code)),
		Error:  err.Error(),
	}
	ctx.JSON(200, r)
}
