package api

import (
	"fmt"
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
	req := dto.ItemUpdateReq{}
	service := handlers.ItemHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.JSON).
		MakeService(&service.Handler).
		Errors

	if err != nil {
		e.Error(500, err, "")
		return
	}

	if err != nil {
		e.Error(500, err, "")
		return
	}

	err = service.Update(&req).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}

	e.OK(req.ID, "ok")

}


func (e Item) Upload(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["Files[]"]

	fmt.Printf("", files)

	req := dto.ItemRescUploadReq{}
	service := handlers.ItemHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.Form, binding.FormMultipart).
		MakeService(&service.Handler).
		Errors


	if err != nil {
		e.Error(500, err, "")
		return
	}

}