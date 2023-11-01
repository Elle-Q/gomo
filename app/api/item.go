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
func (e Item) GetItemAndFilesByItemId(ctx *gin.Context) {
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

	itemVO := vo.ItemWithFilesVO{}
	err = itemService.GetItemAndFilesByItemId(req.ID, &itemVO).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}
	e.OK(itemVO, "ok")
}

// 根据itemId查询文件明细
func (e Item) GetChapter(ctx *gin.Context) {
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

	chapterVO := vo.ChapterVO{}

	err = itemService.GetChapter(req.ID, &chapterVO).Error
	if err != nil {
		e.Error(500, err, "")
		return
	}
	e.OK(chapterVO, "ok")
}
