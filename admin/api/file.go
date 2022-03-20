package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	service "gomo/admin/service"
	"gomo/admin/service/dto"
	"gomo/common/apis"
	"gomo/db/handlers"
)

type File struct {
	apis.Api
}


func (e File) DeleteQNFile(ctx *gin.Context) {
	req := dto.ItemFileDelReq{}
	fileService := service.FileService{
		ItemHandler: &handlers.ItemHandler{},
		FileHandler: &handlers.FileHandler{},
	}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req,binding.JSON).
		MakeService(&fileService.FileHandler.Handler).
		Errors
	if err != nil {
		e.Error(500, err, "")
		return
	}
	err = fileService.DeleteFile(req).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}
	e.OK("", "ok")

}