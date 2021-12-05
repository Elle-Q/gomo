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
	err = service.Find(&req, &user).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(user, "ok")

}