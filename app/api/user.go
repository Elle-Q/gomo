package api

import (
	"github.com/gin-gonic/gin"
	"gomo/app/service/dto"
	"gomo/common/apis"
	"gomo/db/handlers"
	"gomo/db/models"
)

type User struct {
	apis.Api
}

//查询用户信息
func (e User) GetUser(ctx *gin.Context) {
	req := dto.UserApiReq{}
	service := handlers.UserHandler{}
	err := e.MakeContext(ctx).
			MakeDB().
			Bind(&req, nil).
			MakeService(&service.Handler).
			Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}


	var user models.User
	err = service.FindById(&req, &user).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(user, "ok")

}

// 用户信息编辑
func (e User) UpdateUser(ctx *gin.Context) {

	req := dto.UserUpdateApiReq{}
	service := handlers.UserHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, nil).
		MakeService(&service.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	user := req.Generate()
	err = service.Update(&user).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(user, "ok")
}

//登录
func (e User) Login(ctx *gin.Context) {
	req := dto.UserLoginApiReq{}
	service := handlers.UserHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, nil).
		MakeService(&service.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	err = service.Login(&req).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(user, "ok")

}