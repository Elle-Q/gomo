package api

import (
	"github.com/gin-gonic/gin"
	"leetroll/app/service"
	"leetroll/app/service/dto"
	"leetroll/app/service/vo"
	"leetroll/common/apis"
	"leetroll/db/handlers"
	"leetroll/db/models"
)

type Category struct {
	apis.Api
}

func (e Category) List(ctx *gin.Context) {
	catHandler := handlers.CatHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		MakeService(&catHandler.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.Category, 0)

	err = catHandler.List(&list).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(list, "ok")

}

func (e Category) Get(ctx *gin.Context) {
	req := dto.CatApiReq{}
	catHandler := handlers.CatHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, nil).
		MakeService(&catHandler.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	category := models.Category{}

	err = catHandler.Get(req.GetId(), &category).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(category, "ok")
}

func (e Category) ListCatsWithItems(ctx *gin.Context) {
	catItemService := service.NewCatItemService()
	err := e.MakeContext(ctx).
		MakeDB().
		MakeService(&catItemService.ItemHandler.Handler).
		MakeService(&catItemService.CatHandler.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	vos := make([]vo.SubjectVO, 0)

	err = catItemService.ListCatsWithItems(&vos).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(vos, "ok")
}
