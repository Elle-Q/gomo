package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	service "leetroll/admin/service"
	"leetroll/admin/service/dto"
	"leetroll/common/apis"
	"leetroll/db/handlers"
)

type File struct {
	apis.Api
}

// 删除文件  (七牛文件以及数据库记录)
func (e File) DeleteFile(ctx *gin.Context) {
	req := dto.ItemFileDelReq{}
	fileService := service.FileService{
		ItemHandler: &handlers.ItemHandler{},
		FileHandler: &handlers.FileHandler{},
	}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.JSON).
		MakeService(&fileService.FileHandler.Handler).
		Errors
	if err != nil {
		e.Error(500, err, "")
		return
	}
	err = fileService.DeleteFile(req.FileId, req.Bucket, req.Key).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}
	e.OK("", "ok")

}
