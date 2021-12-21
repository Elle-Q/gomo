package api

import (
	"github.com/gin-gonic/gin"
	"gomo/admin/service"
	"gomo/app/service/dto"
	"gomo/common/apis"
	"gomo/db/models"
)

type User struct {
	apis.Api
}

func (e User) GetUser(ctx *gin.Context) {
	req := dto.UserApiReq{}
	service := service.UserService{}
	err := e.MakeContext(ctx).
			MakeDB().
			Bind(&req, nil).
			MakeService(&service.UserHandler.Handler).
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

func (e User) List(ctx *gin.Context) {
	service := service.UserService{}
	err := e.MakeContext(ctx).
		MakeDB().
		MakeService(&service.UserHandler.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.User, 0)

	err = service.List(&list).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(list, "ok")

}