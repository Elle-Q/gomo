package api

import (
	"encoding/json"
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
	err = service.FindById(req.Id, &user).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	userJson,_ := json.Marshal(&user)
	e.OK(userJson, "ok")

}

func (e User) List(ctx *gin.Context) {
	service := handlers.UserHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		MakeService(&service.Handler).
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