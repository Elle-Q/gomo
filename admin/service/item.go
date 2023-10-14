package service

import (
	"leetroll/admin/service/vo"
	"leetroll/db/handlers"
	"leetroll/db/models"
	"leetroll/qiniu"
	"strings"
)

type ItemService struct {
	ItemHandler *handlers.ItemHandler
	FileHandler *handlers.FileHandler
	Error       error
}

func (e *ItemService) GetFilesByItemId(ID int, vo *vo.ItemFilesVO) *ItemService {
	itemHandler := e.ItemHandler
	fileHandler := e.FileHandler

	//获取item的相关信息(名称)
	item := models.MakeItem()
	err := itemHandler.Get(ID, item).Error
	if err != nil {
		e.Error = err
		return e
	}

	files := make([]models.File, 0)
	//获取文件信息
	fileHandler.QueryByItemId(ID, &files)

	main := make([]models.File, 0)
	prev := make([]models.File, 0)
	refs := make([]models.File, 0)

	for _, f := range files {
		p := &f
		if strings.Compare(f.Type, "video") == 0 {
			p.QnLink = qiniu.GetPrivateUrlForM3U8(f.Key)
		} else {
			p.QnLink = qiniu.GetPrivateUrl(f.Key)
		}
		switch f.Type {
		case "main":
			main = append(main, f)
			break
		case "preview":
			prev = append(prev, f)
			break
		case "refs":
			refs = append(refs, f)
			break
		}
	}
	vo.Main = main
	vo.Preview = prev
	vo.Refs = refs
	vo.ID = int64(ID)
	vo.ItemName = item.Name
	vo.RescType = item.Type

	return e
}
