package service

import (
	"leetroll/admin/service/dto"
	"leetroll/db/handlers"
	"leetroll/qiniu"
)

type FileService struct {
	ItemHandler *handlers.ItemHandler
	FileHandler *handlers.FileHandler
	Error       error
}

func (e *FileService) DeleteFile(req dto.ItemFileDelReq) *FileService {
	fileHandler := e.FileHandler

	//删除七牛文件
	err := qiniu.DeleteFile(req.Bucket, req.Key)
	if err != nil {
		e.Error = err
		return e
	}
	//删除数据库记录
	err = fileHandler.Delete(req.FileId).Error
	e.Error = err
	return e
}
