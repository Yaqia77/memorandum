package handler

import (
	"context"

	"github.com/yaqia77/memorandum/apps/api-gateway/internal/service"
	"github.com/yaqia77/memorandum/pkg/e"
	"github.com/yaqia77/memorandum/pkg/res"
	"github.com/yaqia77/memorandum/pkg/util/jwt"

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
