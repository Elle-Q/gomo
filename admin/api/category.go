package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gomo/admin/service/dto"
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


func (e Category) Save(ctx *gin.Context) {
	req := dto.CatUpdateReq{}
	service := handlers.CatHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.Form, binding.FormMultipart).
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

//删除分类
func (e Category) Delete(ctx *gin.Context) {
	req := dto.CatDeleteApiReq{}
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

	err = service.Delete(req.GetId()).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(req.GetId(),"ok")
}
