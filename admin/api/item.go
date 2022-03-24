package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	service "gomo/admin/service"
	"gomo/admin/service/dto"
	"gomo/admin/service/vo"
	"gomo/common/apis"
	"gomo/config"
	"gomo/db/handlers"
	"gomo/db/models"
	"gomo/qiniu"
	"gomo/tool"
)

type Item struct {
	apis.Api
}

//根据itemId查询iten明细
func (e Item) Get(ctx *gin.Context) {
	req := dto.ItemIDReq{}
	handler := handlers.ItemHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, nil).
		MakeService(&handler.Handler).
		Errors
	if err != nil {
		e.Error(500, err, "")
		return
	}

	item := models.MakeItem()
	err = handler.Get(req.ID, item).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}
	e.OK(item, "ok")
}

//根据itemId查询文件明细
func (e Item) GetFilesByItemId(ctx *gin.Context) {
	req := dto.ItemIDReq{}
	itemService := service.NewItemService()
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, nil).
		MakeService(&itemService.ItemHandler.Handler).
		MakeService(&itemService.FileHandler.Handler).
		Errors
	if err != nil {
		e.Error(500, err, "")
		return
	}

	itemVO := vo.ItemFilesVO{}
	err = itemService.GetFilesByItemId(req.ID, &itemVO).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}
	e.OK(itemVO, "ok")
}

func (e Item) List(ctx *gin.Context) {
	handler := handlers.ItemHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		MakeService(&handler.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.Item, 0)

	err = handler.List(&list).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	e.OK(list, "ok")
}

func (e Item) Update(ctx *gin.Context) {
	req := dto.ItemUpdateReq{}
	handler := handlers.ItemHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.JSON).
		MakeService(&handler.Handler).
		Errors
	if err != nil {
		e.Error(500, err, "")
		return
	}
	err = handler.Update(&req).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}
	e.OK(req.ID, "ok")

}

func (e Item) Delete(ctx *gin.Context) {
	req := dto.ItemIDReq{}
	handler := handlers.ItemHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.JSON).
		MakeService(&handler.Handler).
		Errors
	if err != nil {
		e.Error(500, err, "")
		return
	}
	err = handler.Delete(int64(req.ID)).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}
	e.OK(req.ID, "ok")

}


//item 资源文件上传
func (e Item) Upload(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["Files[]"]

	req := dto.ItemRescUploadReq{}
	itemHandler := handlers.ItemHandler{}
	fileHandler := handlers.FileHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.Form).
		MakeService(&itemHandler.Handler).
		MakeService(&fileHandler.Handler).
		Errors
	if err != nil {
		e.Error(500, err, "")
		return
	}

	//上传文件到七牛
	//保存上传进度persistId
	//保存所有信息到db
	for _,fileHeader := range files {
		fileHeader := fileHeader
		go func() {
			key, m3u8Key, e:= qiniu.UploadItemResc(fileHeader, req.Type, req.ItemID)
			if e != nil {
				return
			}
			file := models.File{}
			if len(m3u8Key) > 0 {
				file.QnLink = qiniu.GetPrivateUrlForM3U8(m3u8Key)
				file.Key = m3u8Key
			} else {
				file.QnLink = qiniu.GetPrivateUrl(key)
				file.Key = key
			}
			name,format :=tool.ParseFileName(fileHeader.Filename)
			file.Size = float32(fileHeader.Size)
			file.Format = format
			file.Type = req.Type
			file.Name = name
			file.Bucket = config.QiniuConfig.VideoBucket
			file.ItemId = int64(req.ItemID)
			//保存文件
			errDb := fileHandler.Save(&file).Error

			if errDb != nil {
				return
			}

			fmt.Println("保存文件信息到数据库")
		}()

	}

	e.OK("", "ok")
}
