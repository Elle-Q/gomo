package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"leetroll/admin/service"
	"leetroll/admin/service/dto"
	"leetroll/common/apis"
	"leetroll/db/handlers"
	"leetroll/db/models"
)

type Chapter struct {
	apis.Api
}

// 上传章节信息（目前只支持一章上传）
func (c Chapter) Upload(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	mainFiles := form.File["Main[]"]
	episodesFiles := form.File["Episodes[]"]

	req := dto.ItemChapterUploadReq{}
	fileHandler := handlers.FileHandler{}
	chapterHandler := handlers.ChapterHandler{}
	err := c.MakeContext(ctx).
		MakeDB().
		Bind(&req, binding.Form).
		MakeService(&fileHandler.Handler).
		MakeService(&chapterHandler.Handler).
		Errors
	if err != nil {
		c.Error(500, err, "")
		return
	}

	mainFileIds := uploadByType(req.ItemID, "chapterMain", mainFiles, fileHandler)
	episodesFileIds := uploadByType(req.ItemID, "episodes", episodesFiles, fileHandler)

	chapter := models.Chapter{}
	chapter.ID = int64(req.ChapterID)
	chapter.ItemId = int64(req.ItemID)
	if mainFiles != nil {
		chapter.Main = mainFileIds[0]
	}
	chapterHandler.Save(&chapter)
	chapterHandler.SaveChapterEpisode(int(chapter.ID), episodesFileIds)

	c.OK("", "ok")
}

// 刪除章节下的目录
func (c Chapter) FileDelete(context *gin.Context) {
	req := dto.ChapterFileDelReq{}
	fileService := service.FileService{
		FileHandler: &handlers.FileHandler{},
	}
	chapterHandler := &handlers.ChapterHandler{}
	err := c.MakeContext(context).
		MakeDB().
		Bind(&req, binding.JSON).
		MakeService(&fileService.FileHandler.Handler).
		MakeService(&chapterHandler.Handler).
		Errors
	if err != nil {
		c.Error(500, err, "")
		return
	}
	err = fileService.DeleteFile(req.FileId, req.Bucket, req.Key).Error
	chapterHandler.DelFileByType(req.ChapterId, req.Type, req.FileId)

	if err != nil {
		c.Error(500, err, "")
		return
	}
	c.OK("", "ok")
}
