package api

import (
	"github.com/gin-gonic/gin"
	"gomo/app/service/dto"
	"gomo/common/apis"
	"gomo/db/handlers"
	"gomo/db/models"
)

type Category struct {
	apis.Api
}

func (e Category) List(ctx *gin.Context) {
	service := handlers.CatHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		MakeService(&service.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.Category, 0)

	err = service.List(&list).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(list, "ok")

}

func (e Category) Get(ctx *gin.Context) {
	req := dto.CatApiReq{}
	service := handlers.CatHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req).
		MakeService(&service.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	category := models.Category{}

	err = service.Get(&req, &category).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(category, "ok")
}
