package service

import (
	"leetroll/db/handlers"
	"leetroll/qiniu"
)

type FileService struct {
	ItemHandler *handlers.ItemHandler
	FileHandler *handlers.FileHandler
	Error       error
}

func (e *FileService) DeleteFile(fileId int, bucket string, key string) *FileService {
	fileHandler := e.FileHandler

	//删除七牛文件
	err := qiniu.DeleteFile(bucket, key)
	if err != nil {
		e.Error = err
		return e
	}
	//删除数据库记录
	err = fileHandler.Delete(fileId).Error
	e.Error = err
	return e
}
