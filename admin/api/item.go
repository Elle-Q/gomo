package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gomo/admin/service/dto"
	"gomo/common/apis"
	"gomo/db/handlers"
	"gomo/db/models"
)

type Item struct {
	apis.Api
}

func (e Item) List(ctx *gin.Context) {
	service := handlers.ItemHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		MakeService(&service.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.Item, 0)

	err = service.List(&list).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(list, "ok")
}

func (e Item) Update(ctx *gin.Context) {
	req := dto.CatUpdateReq{}
	service := handlers.CatHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.JSON).
		MakeService(&service.Handler).
		Errors

	if err != nil {
		e.Error(500, err, "")
		return
	}

	Cat := models.Category{}
	err =  req.Generate(&Cat)
	if err != nil {
		e.Error(500, err, "")
		return
	}

	err = service.Save(&Cat).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}

	e.OK(Cat.ID, "ok")

}
