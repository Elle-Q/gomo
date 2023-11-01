package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	service "leetroll/admin/service"
	"leetroll/admin/service/dto"
	"leetroll/admin/service/vo"
	"leetroll/common/apis"
	"leetroll/config"
	"leetroll/db/handlers"
	"leetroll/db/models"
	"leetroll/qiniu"
	"leetroll/tool"
	"mime/multipart"
	"sync"
)

var wg sync.WaitGroup

type Item struct {
	apis.Api
}

// 根据itemId查询iten明细
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

// 根据itemId查询文件明细
func (e Item) GetFilesByItemId(ctx *gin.Context) {
	req := dto.ItemIDReq{}
	itemService := service.NewItemService()
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, nil).
		MakeService(&itemService.ItemHandler.Handler).
		MakeService(&itemService.FileHandler.Handler).
		MakeService(&itemService.ChapterHandler.Handler).
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

// 获取所有的item
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

// 新增或更新item
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

// item 资源文件上传 (main, preview, attachment)
func (e Item) Upload(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	mainFiles := form.File["Main[]"]
	previewFiles := form.File["Preview[]"]
	attachmentFiles := form.File["Attachment[]"]

	req := dto.ItemRescUploadReq{}
	fileHandler := handlers.FileHandler{}
	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.Form).
		MakeService(&fileHandler.Handler).
		Errors
	if err != nil {
		e.Error(500, err, "")
		return
	}

	uploadByType(req.ItemID, "main", mainFiles, fileHandler)
	uploadByType(req.ItemID, "preview", previewFiles, fileHandler)
	uploadByType(req.ItemID, "attachment", attachmentFiles, fileHandler)

	e.OK("", "ok")
}

// 按类型上传 (main, preview, attachment)
// 上传文件到七牛
// 保存上传进度persistId
// 保存所有信息到db
func uploadByType(itemID int, upType string, files []*multipart.FileHeader, handler handlers.FileHandler) []int64 {
	fileIds := make([]int64, 0)
	for _, fileHeader := range files {
		wg.Add(1)
		fileHeader := fileHeader
		go func() {
			defer wg.Done()

			key, m3u8Key, e := qiniu.UploadItemResc(fileHeader, upType, itemID)
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
			name, format := tool.ParseFileName(fileHeader.Filename)
			file.Size = float32(fileHeader.Size)
			file.Format = format
			file.Type = upType
			file.Name = name
			file.Bucket = config.QiniuConfig.VideoBucket
			file.ItemId = int64(itemID)
			//保存文件
			errDb := handler.Save(&file).Error

			fileIds = append(fileIds, file.ID)
			if errDb != nil {
				return
			}
			fmt.Println("保存文件信息到数据库")
		}()
	}
	wg.Wait()
	return fileIds
}
